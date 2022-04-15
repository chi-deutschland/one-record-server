package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/chi-deutschland/one-record-server/pkg/builder"
	"github.com/chi-deutschland/one-record-server/pkg/handler"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	"github.com/chi-deutschland/one-record-server/pkg/service/gcp"
	"github.com/chi-deutschland/one-record-server/pkg/transport/http/middleware"
	"github.com/chi-deutschland/one-record-server/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

var envVars service.Env

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	if err := env.Parse(&envVars); err != nil {
		logrus.Panicf("can`t load .env config: %s", err)
	}
}

func main() {
	secretManager, err := gcp.NewSecretManagerService()
	if err != nil {
		logrus.Panicf("can`t initialize GCP Secret Manager service: %s", err)
	}

	dbService, err := gcp.NewFirestoreService()
	if err != nil {
		logrus.Panicf("can`t initialize GCP Firestore service: %s", err)
	}

	svc := builder.NewServiceBuilder().
		WithEnv(envVars).
		WithGcpSecretManager(secretManager).
		WithGcpFirestore(dbService).
		Build()

	headerAuth, err := utils.NewAuthHeaderSecretValues(svc)
	if err != nil {
		logrus.Panicf("can`t fetch Secrets from Secret Manager: %s", err)
	}

	logrus.WithFields(logrus.Fields{
		"role": svc.Env.ServerRole,
	}).Info("Server will start at :8080")

	fs := http.FileServer(http.Dir(svc.Env.Path.Static))
	router := mux.NewRouter()
	router.PathPrefix("/static").
		Handler(http.StripPrefix("/static", fs))

	router.Use(middleware.LogMiddleware(svc.Env.ServerRole))
	router.Use(middleware.AuthHeaderMiddleware(
		headerAuth.Key,
		headerAuth.Value))
	router.Use(mux.CORSMethodMiddleware(router))

	// Define HandlerFunc for all endpoints here
	rootHandler := handler.NewCompaniesHandler(svc)
	router.HandleFunc("/", rootHandler.Handler).
		Methods(http.MethodGet, http.MethodOptions)

	companyHandler := handler.NewCompanyHandler(svc)
	router.HandleFunc("/{company}", companyHandler.Handler).
		Methods(http.MethodGet, http.MethodOptions)

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.WithFields(logrus.Fields{
		"role": svc.Env.ServerRole,
	}).Fatal(srv.ListenAndServe())
}
