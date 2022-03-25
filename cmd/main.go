package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/chi-deutschland/one-record-server/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"text/template"
	"time"
)

var svc service.Service

type PageData struct {
	Title string
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tmp, err := template.ParseFiles(fmt.Sprintf("%s/root.html", svc.Path.Template))
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
		name := r.Header.Get(svc.Auth.NameKey)
		token := r.Header.Get(svc.Auth.TokenKey)
		if name == svc.Auth.NameValue && token == svc.Auth.TokenValue {
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

func init() {
	if err := env.Parse(&svc); err != nil {
		panic(fmt.Sprintf("cant load .env config: %s", err))
	}
	fmt.Printf("%+v\n", svc)
}

func main() {
	log.Println("Server will start at :8080")
	fs := http.FileServer(http.Dir(svc.Path.Static))
	router := mux.NewRouter()
	router.PathPrefix("/static").
		Handler(http.StripPrefix("/static", fs))

	router.HandleFunc("/", RootHandler).
		Methods(http.MethodGet, http.MethodOptions)

	router.Use(customAuthMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
