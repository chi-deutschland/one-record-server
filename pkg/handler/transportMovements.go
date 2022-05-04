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

type TransportMovementsData struct {
	Title   string
	Host    string
	TransportMovements []model.TransportMovement
}

type TransportMovementsHandler struct {
	Service *service.Service
}

func (h *TransportMovementsHandler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json+ld")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})

		logger.Debugln("\nGET TRANSPORT MOVEMENTS")
		split_url := strings.Split(r.URL.Path[1:], "/")
		companyID := split_url[0]
		pieceID := split_url[2]
		transportMovements, err := h.Service.DBService.GetTransportMovements(
		h.Service.Env.ProjectId,
		companyID, pieceID)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			logger.Debugf("Fetched transportMovements: %#v", transportMovements)
			json.NewEncoder(w).Encode(transportMovements)
		}

	case "POST":
		w.Header().Set("Content-Type", "application/json+ld")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nPOST TRANSPORT MOVEMENT")
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
			var transportMovementID, err = h.Service.DBService.AddTransportMovement(
			h.Service.Env.ProjectId, companyID, pieceID, transportMovement)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				json.NewEncoder(w).Encode(map[string]string{"id": transportMovementID})
				w.WriteHeader(http.StatusCreated)
			}
		}
	}
}

func NewTransportMovementsHandler(svc *service.Service) *TransportMovementsHandler {
	return &TransportMovementsHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*TransportMovementsHandler)(nil)
