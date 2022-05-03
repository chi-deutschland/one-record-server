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

type EventsData struct {
	Title   string
	Host    string
	Events []model.Event
}

type EventsHandler struct {
	Service *service.Service
}

func (h *EventsHandler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json+ld")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})

		logger.Debugln("\nGET EVENTS")
		split_url := strings.Split(r.URL.Path[1:], "/")
		companyID := split_url[0]
		pieceID := split_url[2]
		events, err := h.Service.DBService.GetEvents(
		h.Service.Env.ProjectId,
		companyID, pieceID)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			logger.Debugf("Fetched events: %#v", events)
			json.NewEncoder(w).Encode(events)
		}

	case "POST":
		w.Header().Set("Content-Type", "application/json+ld")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nPOST EVENTS")
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
			var eventID, err = h.Service.DBService.AddEvent(
			h.Service.Env.ProjectId, companyID, pieceID, event)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				json.NewEncoder(w).Encode(map[string]string{"id": eventID})
				w.WriteHeader(http.StatusCreated)
			}
		}
	}
}

func NewEventsHandler(svc *service.Service) *EventsHandler {
	return &EventsHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*EventsHandler)(nil)
