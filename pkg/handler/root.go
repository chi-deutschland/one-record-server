package handler

import (
	"fmt"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"html/template"
	"net/http"
)

type PageData struct {
	Title string
}

type RootHandler struct {
	Service *service.Service
}

func (h *RootHandler) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tmp, err := template.ParseFiles(
		fmt.Sprintf("%s/layout.html", h.Service.Env.Path.Template),
		fmt.Sprintf("%s/root.html", h.Service.Env.Path.Template))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = tmp.Execute(w, PageData{Title: "One Record Server"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewRootHandler(svc *service.Service) *RootHandler {
	return &RootHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*RootHandler)(nil)
