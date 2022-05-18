package handler

import (
	"encoding/json"
	"io"
	"net/http"
    "io/ioutil"
	"github.com/Meschkov/jsonld"
	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"github.com/chi-deutschland/one-record-server/pkg/utils/conv"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ShipmentHandler struct {
	Service *service.Service
}

func (h *ShipmentHandler) Handler(w http.ResponseWriter, r *http.Request) {
	colPath, docPath, id := PathSingleEntry(r.URL.Path, "shipments")
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nGET Shipment")
		logger.Infof("Received request with params %#v", r.URL.Path)

		shipment, err := h.Service.DBService.GetShipment(h.Service.Env.ProjectId, h.Service.Env.ServerRole, docPath)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			var data []byte
			if r.Header.Get(("form")) == "expanded" {
				data, err = jsonld.MarshalExpanded(shipment)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			} else if r.Header.Get(("form")) == "compacted" {
				data, err = jsonld.MarshalCompacted(shipment)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
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
		logger.Infoln("\nPOST Shipment")
		logger.Infof("Received request with params %#v", r.URL.Path)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		var shipment model.Shipment
		err = jsonld.UnmarshalCompacted(body, &shipment)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			ID, err := h.Service.DBService.AddShipment(h.Service.Env.ProjectId, h.Service.Env.ServerRole, colPath, id, shipment)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				json.NewEncoder(w).Encode(map[string]string{"id": ID})
				w.WriteHeader(http.StatusCreated)
			}
		}

	case "PATCH":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPATCH Shipment")
		logger.Infof("Received request with params %#v", r.URL.Path)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		var shipment model.Shipment
		err = jsonld.UnmarshalCompacted(body, &shipment)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			err = h.Service.DBService.UpdateShipment(h.Service.Env.ProjectId, h.Service.Env.ServerRole, docPath, utils.ToFirestoreMap(shipment))
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
		logger.Infoln("\nDELETE Shipment")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var body map[string][]string
		err := decoder.Decode(&body)
		switch {
		case err == io.EOF:
			err = h.Service.DBService.DeleteShipment(h.Service.Env.ProjectId, h.Service.Env.ServerRole, docPath)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		case err != nil:
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			if fields, ok := body["fields"]; ok {
				err = h.Service.DBService.DeleteShipmentFields(h.Service.Env.ProjectId, h.Service.Env.ServerRole, docPath, fields)
				if err != nil {
					// TODO render error message with retry option
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}
}

func NewShipmentHandler(svc *service.Service) *ShipmentHandler {
	return &ShipmentHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*ShipmentHandler)(nil)
