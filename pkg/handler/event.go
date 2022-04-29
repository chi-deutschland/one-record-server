package handler

// import (
// 	"fmt"
// 	"github.com/chi-deutschland/one-record-server/pkg/model"
// 	"github.com/chi-deutschland/one-record-server/pkg/service"
// 	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
// 	"github.com/google/uuid"
// 	"github.com/sirupsen/logrus"
// 	"html/template"
// 	"net/http"
// )

// type EventData struct {
// 	Title   string
// 	Host    string
// 	Events []model.Event
// }

// type EventHandler struct {
// 	Service *service.Service
// }

// func (h *EventHandler) Handler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	logger := logrus.WithFields(logrus.Fields{
// 		"role":       h.Service.Env.ServerRole,
// 		"request_id": uuid.New().String(),
// 	})
// 	logger.Infof("Received request with params %#v", r.URL.Path)
// 	tmp, err := template.ParseFiles(
// 		fmt.Sprintf("%s/layout_iframe.html", h.Service.Env.Path.Template),
// 		fmt.Sprintf("%s/events.html", h.Service.Env.Path.Template))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	companyID := r.URL.Path[1:2]
// 	logisticsObjectID := r.URL.Path[2:]
// 	logger.Debug("Try to fetch events from DB")
// 	events, err := h.Service.DBService.GetEvents(
// 		h.Service.Env.ProjectId,
// 		companyID,
// 		logisticsObjectID)
// 	if err != nil {
// 		// TODO render error message with retry option
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// 	logger.Debugf("Fetched events: %#v", events)

// 	err = tmp.Execute(w, EventData{
// 		Title:   "One Record Server",
// 		Host:    h.Service.Env.Host,
// 		Events: events})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func NewEventHandler(svc *service.Service) *EventHandler {
// 	return &EventHandler{Service: svc}
// }

// var _ onerecordhttp.ContextHandler = (*EventHandler)(nil)
