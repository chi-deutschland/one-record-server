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

type TransportMovementData struct {
	Title   string
	Host    string
	TransportMovement model.TransportMovement
}

type TransportMovementHandler struct {
	Service *service.Service
}

func (h *TransportMovementHandler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nGET TRANSPORT MOVEMENT")
		logger.Infof("Received request with params %#v", r.URL.Path)

		split_url := strings.Split(r.URL.Path[1:], "/")
		companyID := split_url[0]
		pieceID := split_url[2]
		transportMovementID := split_url[4]
		transportMovement, err := h.Service.DBService.GetTransportMovement(
			h.Service.Env.ProjectId,
			companyID,
			pieceID,
			transportMovementID)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			logger.Debugf("Fetched transportMovement: %#v", transportMovement)
			json.NewEncoder(w).Encode(transportMovement)
		}

	case "PATCH":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPATCH TRANSPORT MOVEMENT")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var transportMovement model.TransportMovement
		err := decoder.Decode(&transportMovement)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			split_url := strings.Split(r.URL.Path[1:], "/")
			companyID := split_url[0]
			pieceID := split_url[2]
			transportMovementID := split_url[4]
			logger.Debugln(companyID, pieceID, transportMovementID)
			logger.Debugf("transportMovement: %#v", transportMovement)
			h.Service.DBService.UpdateTransportMovement(
			h.Service.Env.ProjectId,
			companyID, pieceID, transportMovementID, transportMovement)
		}

	case "DELETE":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nDELETE TRANSPORT MOVEMENT")
		logger.Infof("Received request with params %#v", r.URL.Path)

		split_url := strings.Split(r.URL.Path[1:], "/")
		companyID := split_url[0]
		pieceID := split_url[2]
		transportMovementID := split_url[4]
		decoder := json.NewDecoder(r.Body)
		var body map[string][]string
		err := decoder.Decode(&body)
		switch {
		case err == io.EOF:
			h.Service.DBService.DeleteTransportMovement(
				h.Service.Env.ProjectId,
				companyID, pieceID, transportMovementID)
		case err != nil:
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			if fields, ok := body["fields"]; ok {
				h.Service.DBService.DeleteTransportMovementFields(
				h.Service.Env.ProjectId,
				companyID, pieceID, transportMovementID, fields)
			}
		}
	}
}

func NewTransportMovementHandler(svc *service.Service) *TransportMovementHandler {
	return &TransportMovementHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*TransportMovementHandler)(nil)
