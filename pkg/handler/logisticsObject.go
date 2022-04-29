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

// type LogisticsObjectData struct {
// 	Title   string
// 	Host    string
// 	LogisticsObject model.LogisticsObject
// }

// type LogisticsObjectHandler struct {
// 	Service *service.Service
// }

// func (h *LogisticsObjectHandler) Handler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	logger := logrus.WithFields(logrus.Fields{
// 		"role":       h.Service.Env.ServerRole,
// 		"request_id": uuid.New().String(),
// 	})
// 	logger.Infof("Received request with params %#v", r.URL.Path)
// 	tmp, err := template.ParseFiles(
// 		fmt.Sprintf("%s/layout_iframe.html", h.Service.Env.Path.Template),
// 		fmt.Sprintf("%s/logisticsObject.html", h.Service.Env.Path.Template))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	companyID := r.URL.Path[1:2]
// 	logisticsObjectID := r.URL.Path[2:]
// 	logger.Debug("Try to fetch logisticsObjects from DB")
// 	logisticsObject, err := h.Service.DBService.GetLogisticsObject(
// 		h.Service.Env.ProjectId,
// 		companyID, logisticsObjectID)
// 	if err != nil {
// 		// TODO render error message with retry option
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// 	logger.Debugf("Fetched logisticsObject: %#v", logisticsObject)

// 	err = tmp.Execute(w, LogisticsObjectData{
// 		Title:   "One Record Server",
// 		Host:    h.Service.Env.Host,
// 		LogisticsObject: logisticsObject})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func NewLogisticsObjectHandler(svc *service.Service) *LogisticsObjectHandler {
// 	return &LogisticsObjectHandler{Service: svc}
// }

// var _ onerecordhttp.ContextHandler = (*LogisticsObjectHandler)(nil)
