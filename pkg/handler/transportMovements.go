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

type TransportMovementsHandler struct {
	Service *service.Service
}

func (h *TransportMovementsHandler) Handler(w http.ResponseWriter, r *http.Request) {
	path := PathMultipleEntries(r.URL.Path)
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json+ld")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nGET TransportMovements")
		logger.Infof("Received request with params %#v", r.URL.Path)

		transportMovements, err := h.Service.DBService.GetTransportMovements(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			json.NewEncoder(w).Encode(transportMovements)
		}

	case "POST":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPOST TransportMovement")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var transportMovement model.TransportMovement
		err := decoder.Decode(&transportMovement)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			ID, err := h.Service.DBService.AddTransportMovement(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, transportMovement.ID, transportMovement)
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

func NewTransportMovementsHandler(svc *service.Service) *TransportMovementsHandler {
	return &TransportMovementsHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*TransportMovementsHandler)(nil)
