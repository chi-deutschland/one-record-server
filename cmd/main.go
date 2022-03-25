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
	w.Header().Set("Content-Type", "text/html")
	tmp, err := template.ParseFiles(fmt.Sprintf("%s/root.html", TemplatePathPrefix))
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

func customAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		name := r.Header.Get(AuthHeaderNameKey)
		token := r.Header.Get(AuthHeaderTokenKey)
		if name == AuthHeaderNameValue && token == AuthHeaderTokenValue {
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

func main() {
	log.Println("Server will start at :8000")
	fs := http.FileServer(http.Dir("./web/public/static/"))
	router := mux.NewRouter()
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", fs))
	router.HandleFunc("/", RootHandler).Methods(http.MethodGet, http.MethodOptions)
	router.Use(customAuthMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))
	log.Fatal(http.ListenAndServe(":8000", router))
}
