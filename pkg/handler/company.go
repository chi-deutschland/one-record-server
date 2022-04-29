package handler

import (
	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
    "encoding/json"
)

type CompanyData struct {
	Title   string
	Host    string
	Company model.Company
}

type CompanyHandler struct {
	Service *service.Service
}

func (h *CompanyHandler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("GET COMPANY")
		logger.Infof("Received request with params %#v", r.URL.Path)

		companyID := r.URL.Path[1:]
		logger.Debug("Try to fetch companies from DB")
		company, err := h.Service.DBService.GetCompany(
			h.Service.Env.ProjectId,
			companyID)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			logger.Debugf("Fetched company: %#v", company)
			json.NewEncoder(w).Encode(company)
		}

	case "PATCH":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("PATCH COMPANY")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var company model.Company
		err := decoder.Decode(&company)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			companyID := r.URL.Path[1:]
			logger.Debugln(companyID)
			logger.Debugf("company: %#v", company)
			h.Service.DBService.UpdateCompany(
			h.Service.Env.ProjectId,
			companyID, company)
		}

	case "DELETE":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("DELETE COMPANY")
		logger.Infof("Received request with params %#v", r.URL.Path)

		companyID := r.URL.Path[1:]
		h.Service.DBService.DeleteCompany(
			h.Service.Env.ProjectId,
			companyID)
	}
}

func NewCompanyHandler(svc *service.Service) *CompanyHandler {
	return &CompanyHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*CompanyHandler)(nil)
