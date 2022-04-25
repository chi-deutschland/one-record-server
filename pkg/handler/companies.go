package handler

import (
	"fmt"
	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
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
	w.Header().Set("Content-Type", "text/html")
	logger := logrus.WithFields(logrus.Fields{
		"role":       h.Service.Env.ServerRole,
		"request_id": uuid.New().String(),
	})

	tmp, err := template.ParseFiles(
		fmt.Sprintf("%s/layout.html", h.Service.Env.Path.Template),
		fmt.Sprintf("%s/companies.html", h.Service.Env.Path.Template))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	logger.Debug("Try to fetch companies from DB")
	companies, err := h.Service.DBService.GetCompanies(h.Service.Env.ProjectId, h.Service.Env.ServerRole)
	if err != nil {
		// TODO render error message with retry option
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	logger.Debugf("Fetched companies: %#v", companies)
	pageData := CompaniesData{
		Title:     "One Record Server - Companies",
		Host:      h.Service.Env.Host,
		Companies: companies}

	err = tmp.Execute(w, pageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewCompaniesHandler(svc *service.Service) *CompaniesHandler {
	return &CompaniesHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*CompaniesHandler)(nil)
