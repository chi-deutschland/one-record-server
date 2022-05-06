package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"github.com/chi-deutschland/one-record-server/pkg/utils/conv"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type EventHandler struct {
	Service *service.Service
}

func (h *EventHandler) Handler(w http.ResponseWriter, r *http.Request) {
	path := PathMultipleEntries(r.URL.Path)
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nGET Event")
		logger.Infof("Received request with params %#v", r.URL.Path)

		event, err := h.Service.DBService.GetEvent(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			json.NewEncoder(w).Encode(event)
		}

	case "PATCH":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPATCH Event")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var event model.Event
		err := decoder.Decode(&event)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			err = h.Service.DBService.UpdateEvent(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, utils.ToFirestoreMap(event))
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

	case "DELETE":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nDELETE Event")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var body map[string][]string
		err := decoder.Decode(&body)
		switch {
		case err == io.EOF:
			err = h.Service.DBService.DeleteEvent(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		case err != nil:
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			if fields, ok := body["fields"]; ok {
				err = h.Service.DBService.DeleteEventFields(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, fields)
				if err != nil {
					// TODO render error message with retry option
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}
}

func NewEventHandler(svc *service.Service) *EventHandler {
	return &EventHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*EventHandler)(nil)
