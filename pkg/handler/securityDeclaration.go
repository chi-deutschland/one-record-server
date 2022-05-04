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

type SecurityDeclarationData struct {
	Title   string
	Host    string
	SecurityDeclaration model.SecurityDeclaration
}

type SecurityDeclarationHandler struct {
	Service *service.Service
}

func (h *SecurityDeclarationHandler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nPOST SECURITY DECLARATION")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var securityDeclaration model.SecurityDeclaration
		err := decoder.Decode(&securityDeclaration)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			split_url := strings.Split(r.URL.Path[1:], "/")
			companyID := split_url[0]
			pieceID := split_url[2]
			err = h.Service.DBService.AddSecurityDeclaration(
			h.Service.Env.ProjectId, companyID, pieceID, securityDeclaration)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusCreated)
			}
		}
	
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nGET SECURITY DECLARATION")
		logger.Infof("Received request with params %#v", r.URL.Path)

		split_url := strings.Split(r.URL.Path[1:], "/")
		companyID := split_url[0]
		pieceID := split_url[2]
		securityDeclaration, err := h.Service.DBService.GetSecurityDeclaration(
			h.Service.Env.ProjectId,
			companyID,
			pieceID)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			logger.Debugf("Fetched securityDeclaration: %#v", securityDeclaration)
			json.NewEncoder(w).Encode(securityDeclaration)
		}

	case "PATCH":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPATCH SECURITY DECLARATION")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var securityDeclaration model.SecurityDeclaration
		err := decoder.Decode(&securityDeclaration)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			split_url := strings.Split(r.URL.Path[1:], "/")
			companyID := split_url[0]
			pieceID := split_url[2]
			logger.Debugln(companyID, pieceID)
			logger.Debugf("securityDeclaration: %#v", securityDeclaration)
			h.Service.DBService.UpdateSecurityDeclaration(
			h.Service.Env.ProjectId,
			companyID, pieceID, securityDeclaration)
		}

	case "DELETE":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nDELETE SECURITY DECLARATION")
		logger.Infof("Received request with params %#v", r.URL.Path)

		split_url := strings.Split(r.URL.Path[1:], "/")
		companyID := split_url[0]
		pieceID := split_url[2]
		decoder := json.NewDecoder(r.Body)
		var body map[string][]string
		err := decoder.Decode(&body)
		switch {
		case err == io.EOF:
			h.Service.DBService.DeleteSecurityDeclaration(
				h.Service.Env.ProjectId,
				companyID, pieceID)
		case err != nil:
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			if fields, ok := body["fields"]; ok {
				h.Service.DBService.DeleteSecurityDeclarationFields(
				h.Service.Env.ProjectId,
				companyID, pieceID, fields)
			}
		}
	}
}

func NewSecurityDeclarationHandler(svc *service.Service) *SecurityDeclarationHandler {
	return &SecurityDeclarationHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*SecurityDeclarationHandler)(nil)