package handler

import (
	"encoding/json"
	"net/http"

	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ExternalReferencesHandler struct {
	Service *service.Service
}

func (h *ExternalReferencesHandler) Handler(w http.ResponseWriter, r *http.Request) {
	path := PathMultipleEntries(r.URL.Path)
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json+ld")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nGET ExternalReferences")
		logger.Infof("Received request with params %#v", r.URL.Path)

		externalReferences, err := h.Service.DBService.GetExternalReferences(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			json.NewEncoder(w).Encode(externalReferences)
		}

	case "POST":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPOST ExternalReference")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var externalReference model.ExternalReference
		err := decoder.Decode(&externalReference)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			ID, err := h.Service.DBService.AddExternalReference(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, externalReference.ID, externalReference)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				json.NewEncoder(w).Encode(map[string]string{"id": ID})
				w.WriteHeader(http.StatusCreated)
			}
		}
	}
}

func NewExternalReferencesHandler(svc *service.Service) *ExternalReferencesHandler {
	return &ExternalReferencesHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*ExternalReferencesHandler)(nil)
