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

type CompaniesData struct {
	Title     string
	Host      string
	Companies []model.Company
}

type CompaniesHandler struct {
	Service *service.Service
}

func (h *CompaniesHandler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json+ld")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})

		logger.Debugln("\nGET COMPANIES")
		companies, err := h.Service.DBService.GetCompanies(h.Service.Env.ProjectId, h.Service.Env.ServerRole)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			logger.Debugf("Fetched companies: %#v", companies)
			json.NewEncoder(w).Encode(companies)
		}

	case "POST":
		w.Header().Set("Content-Type", "application/json+ld")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPOST COMPANY")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var company model.Company
		err := decoder.Decode(&company)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			var companyID, err = h.Service.DBService.AddCompany(
			h.Service.Env.ProjectId,
			company)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				json.NewEncoder(w).Encode(map[string]string{"id": companyID})
				w.WriteHeader(http.StatusCreated)
			}
		}
	}
}

func NewCompaniesHandler(svc *service.Service) *CompaniesHandler {
	return &CompaniesHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*CompaniesHandler)(nil)
