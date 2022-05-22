package handler

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/chi-deutschland/one-record-server/pkg/jsonld"
	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type RegulatedEntitiesHandler struct {
	Service *service.Service
}

func (h *RegulatedEntitiesHandler) Handler(w http.ResponseWriter, r *http.Request) {
	path := PathMultipleEntries(r.URL.Path)
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json+ld")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nGET RegulatedEntities")
		logger.Infof("Received request with params %#v", r.URL.Path)

		regulatedEntities, err := h.Service.DBService.GetRegulatedEntities(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			data, err := jsonld.MarshalCompacted(regulatedEntities)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			_, err = w.Write(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

	case "POST":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPOST RegulatedEntity")
		logger.Infof("Received request with params %#v", r.URL.Path)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		var regulatedEntity model.RegulatedEntity
		err = jsonld.UnmarshalCompacted(body, &regulatedEntity)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			ID, err := h.Service.DBService.AddRegulatedEntity(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, regulatedEntity.ID, regulatedEntity)
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

func NewRegulatedEntitiesHandler(svc *service.Service) *RegulatedEntitiesHandler {
	return &RegulatedEntitiesHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*RegulatedEntitiesHandler)(nil)