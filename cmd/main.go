package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"text/template"
)

var TemplatePathPrefix = "web/public/templates"

type PageData struct {
	Title string
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")
		tmp, err := template.ParseFiles(fmt.Sprintf("%s/root.gohtml", TemplatePathPrefix))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmp.Execute(w, PageData{Title: "One Record Server"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	log.Println("Server will start at")
	fs := http.FileServer(http.Dir("./web/public/static/"))
	route := mux.NewRouter()
	route.PathPrefix("/static").Handler(http.StripPrefix("/static", fs))
	route.HandleFunc("/", RootHandler)
	log.Fatal(http.ListenAndServe(":8000", route))
}
