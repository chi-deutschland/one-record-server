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

type PageData struct {
	Title     string
	Companies []model.Company
}

type RootHandler struct {
	Service *service.Service
}

func (h *RootHandler) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	logger := logrus.WithFields(logrus.Fields{
		"role":       h.Service.Env.ServerRole,
		"request_id": uuid.New().String(),
	})

	tmp, err := template.ParseFiles(
		fmt.Sprintf("%s/layout.html", h.Service.Env.Path.Template),
		fmt.Sprintf("%s/root.html", h.Service.Env.Path.Template))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	logger.Debug("Try to fetch companies from DB")
	companies, err := h.Service.DBService.GetCompanies(h.Service.Env.ProjectId)
	if err != nil {
		// TODO render error message with retry option
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	logger.Debugf("Fetched companies: %#v", companies)
	pageData := PageData{Title: "One Record Server - Companies", Companies: companies}

	err = tmp.Execute(w, pageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewRootHandler(svc *service.Service) *RootHandler {
	return &RootHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*RootHandler)(nil)
