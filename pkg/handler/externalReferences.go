package handler

import (
	"net/http"
	"strings"
	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
    "encoding/json"
)

type ExternalReferencesData struct {
	Title   string
	Host    string
	ExternalReferences []model.ExternalReference
}

type ExternalReferencesHandler struct {
	Service *service.Service
}

func (h *ExternalReferencesHandler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json+ld")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})

		logger.Debugln("\nGET EXTERNAL REFERENCES")
		split_url := strings.Split(r.URL.Path[1:], "/")
		companyID := split_url[0]
		pieceID := split_url[2]
		externalReferences, err := h.Service.DBService.GetExternalReferences(
		h.Service.Env.ProjectId,
		companyID, pieceID)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			logger.Debugf("Fetched externalReferences: %#v", externalReferences)
			json.NewEncoder(w).Encode(externalReferences)
		}

	case "POST":
		w.Header().Set("Content-Type", "application/json+ld")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nPOST EXTERNAL REFERENCE")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var externalReference model.ExternalReference
		err := decoder.Decode(&externalReference)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			split_url := strings.Split(r.URL.Path[1:], "/")
			companyID := split_url[0]
			pieceID := split_url[2]
			var externalReferenceID, err = h.Service.DBService.AddExternalReference(
			h.Service.Env.ProjectId, companyID, pieceID, externalReference)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				json.NewEncoder(w).Encode(map[string]string{"id": externalReferenceID})
				w.WriteHeader(http.StatusCreated)
			}
		}
	}
}

func NewExternalReferencesHandler(svc *service.Service) *ExternalReferencesHandler {
	return &ExternalReferencesHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*ExternalReferencesHandler)(nil)
