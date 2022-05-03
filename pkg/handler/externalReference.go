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

type ExternalReferenceData struct {
	Title   string
	Host    string
	ExternalReference model.ExternalReference
}

type ExternalReferenceHandler struct {
	Service *service.Service
}

func (h *ExternalReferenceHandler) Handler(w http.ResponseWriter, r *http.Request) {
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
		externalReferenceID := split_url[4]
		externalReference, err := h.Service.DBService.GetExternalReference(
			h.Service.Env.ProjectId,
			companyID,
			pieceID,
			externalReferenceID)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			logger.Debugf("Fetched externalReference: %#v", externalReference)
			json.NewEncoder(w).Encode(externalReference)
		}

	case "PATCH":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPATCH EVENT")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var externalReference model.ExternalReference
		err := decoder.Decode(&externalReference)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			split_url := strings.Split(r.URL.Path[1:], "/")
			companyID := split_url[0]
			pieceID := split_url[2]
			externalReferenceID := split_url[4]
			logger.Debugln(companyID, pieceID, externalReferenceID)
			logger.Debugf("externalReference: %#v", externalReference)
			h.Service.DBService.UpdateExternalReference(
			h.Service.Env.ProjectId,
			companyID, pieceID, externalReferenceID, externalReference)
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
		externalReferenceID := split_url[4]
		decoder := json.NewDecoder(r.Body)
		var body map[string][]string
		err := decoder.Decode(&body)
		switch {
		case err == io.EOF:
			h.Service.DBService.DeleteExternalReference(
				h.Service.Env.ProjectId,
				companyID, pieceID, externalReferenceID)
		case err != nil:
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			if fields, ok := body["fields"]; ok {
				h.Service.DBService.DeleteExternalReferenceFields(
				h.Service.Env.ProjectId,
				companyID, pieceID, externalReferenceID, fields)
			}
		}
	}
}

func NewExternalReferenceHandler(svc *service.Service) *ExternalReferenceHandler {
	return &ExternalReferenceHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*ExternalReferenceHandler)(nil)
