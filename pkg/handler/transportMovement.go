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

type TransportMovementHandler struct {
	Service *service.Service
}

func (h *TransportMovementHandler) Handler(w http.ResponseWriter, r *http.Request) {
	path := PathMultipleEntries(r.URL.Path)
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nGET TransportMovement")
		logger.Infof("Received request with params %#v", r.URL.Path)

		transportMovement, err := h.Service.DBService.GetTransportMovement(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			json.NewEncoder(w).Encode(transportMovement)
		}

	case "PATCH":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPATCH TransportMovement")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var transportMovement model.TransportMovement
		err := decoder.Decode(&transportMovement)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			err = h.Service.DBService.UpdateTransportMovement(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, utils.ToFirestoreMap(transportMovement))
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
		logger.Infoln("\nDELETE TransportMovement")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var body map[string][]string
		err := decoder.Decode(&body)
		switch {
		case err == io.EOF:
			err = h.Service.DBService.DeleteTransportMovement(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		case err != nil:
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			if fields, ok := body["fields"]; ok {
				err = h.Service.DBService.DeleteTransportMovementFields(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, fields)
				if err != nil {
					// TODO render error message with retry option
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}
}

func NewTransportMovementHandler(svc *service.Service) *TransportMovementHandler {
	return &TransportMovementHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*TransportMovementHandler)(nil)
