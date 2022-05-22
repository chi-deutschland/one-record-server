package handler

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	"github.com/chi-deutschland/one-record-server/pkg/jsonld"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CompaniesHandler struct {
	Service *service.Service
}

func (h *CompaniesHandler) Handler(w http.ResponseWriter, r *http.Request) {
	path := PathMultipleEntries(r.URL.Path)
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json+ld")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nGET Companies")
		logger.Infof("Received request with params %#v", r.URL.Path)

		companies, err := h.Service.DBService.GetCompanies(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			data, err := jsonld.MarshalCompacted(companies)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
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
		logger.Infoln("\nPOST Company")
		logger.Infof("Received request with params %#v", r.URL.Path)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		var company model.Company
		err = jsonld.UnmarshalCompacted(body, &company)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			ID, err := h.Service.DBService.AddCompany(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, company.ID, company)
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

func NewCompaniesHandler(svc *service.Service) *CompaniesHandler {
	return &CompaniesHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*CompaniesHandler)(nil)
