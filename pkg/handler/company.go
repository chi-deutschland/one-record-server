package handler

import (
	"fmt"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

type CompanyData struct {
	Title     string
	CompanyID string
}

type CompanyHandler struct {
	Service *service.Service
}

func (h *CompanyHandler) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	logger := logrus.WithFields(logrus.Fields{
		"role":       h.Service.Env.ServerRole,
		"request_id": uuid.New().String(),
	})
	logger.Infof("Received request with params %#v", r.URL.Path)
	tmp, err := template.ParseFiles(
		fmt.Sprintf("%s/layout_iframe.html", h.Service.Env.Path.Template),
		fmt.Sprintf("%s/company.html", h.Service.Env.Path.Template))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = tmp.Execute(w, PageData{Title: "One Record Server"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewCompanyHandler(svc *service.Service) *CompanyHandler {
	return &CompanyHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*CompanyHandler)(nil)
