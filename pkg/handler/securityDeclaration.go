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

type SecurityDeclarationHandler struct {
	Service *service.Service
}

func (h *SecurityDeclarationHandler) Handler(w http.ResponseWriter, r *http.Request) {
	colPath, docPath, id := PathSingleEntry(r.URL.Path, "securityDeclarations")
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nGET SecurityDeclaration")
		logger.Infof("Received request with params %#v", r.URL.Path)

		securityDeclaration, err := h.Service.DBService.GetSecurityDeclaration(h.Service.Env.ProjectId, h.Service.Env.ServerRole, docPath)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			json.NewEncoder(w).Encode(securityDeclaration)
		}

	case "POST":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPOST SecurityDeclaration")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var securityDeclaration model.SecurityDeclaration
		err := decoder.Decode(&securityDeclaration)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			ID, err := h.Service.DBService.AddSecurityDeclaration(h.Service.Env.ProjectId, h.Service.Env.ServerRole, colPath, id, securityDeclaration)
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
		logger.Infoln("\nPATCH SecurityDeclaration")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var securityDeclaration model.SecurityDeclaration
		err := decoder.Decode(&securityDeclaration)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			err = h.Service.DBService.UpdateSecurityDeclaration(h.Service.Env.ProjectId, h.Service.Env.ServerRole, docPath, utils.ToFirestoreMap(securityDeclaration))
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
		logger.Infoln("\nDELETE SecurityDeclaration")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var body map[string][]string
		err := decoder.Decode(&body)
		switch {
		case err == io.EOF:
			err = h.Service.DBService.DeleteSecurityDeclaration(h.Service.Env.ProjectId, h.Service.Env.ServerRole, docPath)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		case err != nil:
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			if fields, ok := body["fields"]; ok {
				err = h.Service.DBService.DeleteSecurityDeclarationFields(h.Service.Env.ProjectId, h.Service.Env.ServerRole, docPath, fields)
				if err != nil {
					// TODO render error message with retry option
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}
}

func NewSecurityDeclarationHandler(svc *service.Service) *SecurityDeclarationHandler {
	return &SecurityDeclarationHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*SecurityDeclarationHandler)(nil)
