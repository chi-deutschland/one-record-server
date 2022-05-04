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
	"io"
)

type EventData struct {
	Title   string
	Host    string
	Event model.Event
}

type EventHandler struct {
	Service *service.Service
}

func (h *EventHandler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nGET EVENTS")
		logger.Infof("Received request with params %#v", r.URL.Path)

		split_url := strings.Split(r.URL.Path[1:], "/")
		companyID := split_url[0]
		pieceID := split_url[2]
		eventID := split_url[4]
		event, err := h.Service.DBService.GetEvent(
			h.Service.Env.ProjectId,
			companyID,
			pieceID,
			eventID)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			logger.Debugf("Fetched event: %#v", event)
			json.NewEncoder(w).Encode(event)
		}

	case "PATCH":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPATCH EVENT")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var event model.Event
		err := decoder.Decode(&event)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			split_url := strings.Split(r.URL.Path[1:], "/")
			companyID := split_url[0]
			pieceID := split_url[2]
			eventID := split_url[4]
			logger.Debugln(companyID, pieceID, eventID)
			logger.Debugf("event: %#v", event)
			h.Service.DBService.UpdateEvent(
			h.Service.Env.ProjectId,
			companyID, pieceID, eventID, event)
		}

	case "DELETE":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nDELETE EVENT")
		logger.Infof("Received request with params %#v", r.URL.Path)

		split_url := strings.Split(r.URL.Path[1:], "/")
		companyID := split_url[0]
		pieceID := split_url[2]
		eventID := split_url[4]
		decoder := json.NewDecoder(r.Body)
		var body map[string][]string
		err := decoder.Decode(&body)
		switch {
		case err == io.EOF:
			h.Service.DBService.DeleteEvent(
				h.Service.Env.ProjectId,
				companyID, pieceID, eventID)
		case err != nil:
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			if fields, ok := body["fields"]; ok {
				h.Service.DBService.DeleteEventFields(
				h.Service.Env.ProjectId,
				companyID, pieceID, eventID, fields)
			}
		}
	}
}

func NewEventHandler(svc *service.Service) *EventHandler {
	return &EventHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*EventHandler)(nil)