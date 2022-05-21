package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/chi-deutschland/one-record-server/pkg/builder"
	"github.com/chi-deutschland/one-record-server/pkg/handler"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	"github.com/chi-deutschland/one-record-server/pkg/service/gcp"
	"github.com/chi-deutschland/one-record-server/pkg/transport/http/middleware"
	"github.com/chi-deutschland/one-record-server/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
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

	fcm, err := gcp.NewFCM()
	if err != nil {
		logrus.Panicf("can`t subscribe: %s", err)
	}

	ps, err := gcp.NewPubSub()
	if err != nil {
		logrus.Panicf("can`t subscribe: %s", err)
	}


	svc := builder.NewServiceBuilder().
		WithEnv(envVars).
		WithGcpSecretManager(secretManager).
		WithGcpFirestore(dbService).
		WithFCM(fcm).
		WithPS(ps).
		Build()


	headerAuth, err := utils.NewAuthHeaderSecretValues(svc)
	if err != nil {
		logrus.Panicf("can`t fetch Secrets from Secret Manager: %s", err)
	}

	logrus.WithFields(logrus.Fields{
		"role": svc.Env.ServerRole,
	}).Info("Server will start at",svc.Env.Host)

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
	companiesHandler := handler.NewCompaniesHandler(svc)
	router.HandleFunc("/companies", companiesHandler.Handler).
	Methods(http.MethodGet, http.MethodPost, http.MethodOptions)

	companyHandler := handler.NewCompanyHandler(svc)
	router.HandleFunc("/companies/{company}", companyHandler.Handler).
	Methods(http.MethodGet, http.MethodPatch, http.MethodDelete, http.MethodOptions)

	piecesHandler := handler.NewPiecesHandler(svc)
	router.HandleFunc("/companies/{company}/pieces", piecesHandler.Handler).
	Methods(http.MethodGet, http.MethodPost, http.MethodOptions)

	multiplePiecesHandler := handler.NewMultiplePiecesHandler(svc)
	router.HandleFunc("/pieces", multiplePiecesHandler.Handler).
	Methods(http.MethodGet, http.MethodOptions)

	pieceHandler := handler.NewPieceHandler(svc)
	router.HandleFunc("/companies/{company}/pieces/{piece}", pieceHandler.Handler).
	Methods(http.MethodGet, http.MethodPatch, http.MethodDelete, http.MethodOptions)

	itemsHandler := handler.NewItemsHandler(svc)
	router.HandleFunc("/companies/{company}/pieces/{piece}/items", itemsHandler.Handler).
	Methods(http.MethodGet, http.MethodPost, http.MethodOptions)

	itemHandler := handler.NewItemHandler(svc)
	router.HandleFunc("/companies/{company}/pieces/{piece}/items/{item}", itemHandler.Handler).
	Methods(http.MethodGet, http.MethodPatch, http.MethodDelete, http.MethodOptions)

	eventsHandler := handler.NewEventsHandler(svc)
	router.HandleFunc("/companies/{company}/pieces/{piece}/events", eventsHandler.Handler).
	Methods(http.MethodGet, http.MethodPost, http.MethodOptions)

	eventHandler := handler.NewEventHandler(svc)
	router.HandleFunc("/companies/{company}/pieces/{piece}/events/{event}", eventHandler.Handler).
	Methods(http.MethodGet, http.MethodPatch, http.MethodDelete, http.MethodOptions)

	externalReferencesHandler := handler.NewExternalReferencesHandler(svc)
	router.HandleFunc("/companies/{company}/pieces/{piece}/externalReferences", externalReferencesHandler.Handler).
	Methods(http.MethodGet, http.MethodPost, http.MethodOptions)

	externalReferenceHandler := handler.NewExternalReferenceHandler(svc)
	router.HandleFunc("/companies/{company}/pieces/{piece}/externalReferences/{externalReference}", externalReferenceHandler.Handler).
	Methods(http.MethodGet, http.MethodPatch, http.MethodDelete, http.MethodOptions)

	transportMovementsHandler := handler.NewTransportMovementsHandler(svc)
	router.HandleFunc("/companies/{company}/pieces/{piece}/transportMovements", transportMovementsHandler.Handler).
	Methods(http.MethodGet, http.MethodPost, http.MethodOptions)

	transportMovementHandler := handler.NewTransportMovementHandler(svc)
	router.HandleFunc("/companies/{company}/pieces/{piece}/transportMovements/{transportMovement}", transportMovementHandler.Handler).
	Methods(http.MethodGet, http.MethodPatch, http.MethodDelete, http.MethodOptions)

	securityDeclarationHandler := handler.NewSecurityDeclarationHandler(svc)
	router.HandleFunc("/companies/{company}/pieces/{piece}/securityDeclaration", securityDeclarationHandler.Handler).
	Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodOptions)

	regulatedEntitiesHandler := handler.NewRegulatedEntitiesHandler(svc)
	router.HandleFunc("/regulatedEntities", regulatedEntitiesHandler.Handler).
	Methods(http.MethodGet, http.MethodPost, http.MethodOptions)

	regulatedEntityHandler := handler.NewRegulatedEntityHandler(svc)
	router.HandleFunc("/regulatedEntities/{regulatedEntity}", regulatedEntityHandler.Handler).
	Methods(http.MethodGet, http.MethodPatch, http.MethodDelete, http.MethodOptions)

	regulatedEntitiesHandlerSecurityDeclaration := handler.NewRegulatedEntitiesHandler(svc)
	router.HandleFunc("/companies/{company}/pieces/{piece}/regulatedEntities", regulatedEntitiesHandlerSecurityDeclaration.Handler).
	Methods(http.MethodGet, http.MethodPost, http.MethodOptions)

	regulatedEntityHandlerSecurityDeclaration := handler.NewRegulatedEntityHandler(svc)
	router.HandleFunc("/companies/{company}/pieces/{piece}/regulatedEntities/{regulatedEntity}", regulatedEntityHandlerSecurityDeclaration.Handler).
	Methods(http.MethodGet, http.MethodPatch, http.MethodDelete, http.MethodOptions)

	shipmentHandler := handler.NewShipmentHandler(svc)
	router.HandleFunc("/companies/{company}/pieces/{piece}/shipment", shipmentHandler.Handler).
	Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodOptions)

	dataHandler := handler.NewDataHandler(svc)
	router.HandleFunc("/data", dataHandler.Handler).
	Methods(http.MethodPost, http.MethodDelete, http.MethodOptions)

	subscriptionHandler := handler.NewSubscriptionHandler(svc)
	router.HandleFunc("/sub", subscriptionHandler.Handler).
	Methods(http.MethodPost, http.MethodOptions)

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println(srv)
	logrus.WithFields(logrus.Fields{
		"role": svc.Env.ServerRole,
	}).Fatal(srv.ListenAndServe())
}