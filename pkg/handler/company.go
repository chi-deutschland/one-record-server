package handler

import (
	"encoding/json"
	"io"
	"net/http"
    "io/ioutil"
	"github.com/Meschkov/jsonld"
	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"github.com/chi-deutschland/one-record-server/pkg/utils/conv"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CompanyHandler struct {
	Service *service.Service
}

func (h *CompanyHandler) Handler(w http.ResponseWriter, r *http.Request) {
	path := PathMultipleEntries(r.URL.Path)
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nGET Company")
		logger.Infof("Received request with params %#v", r.URL.Path)

		company, err := h.Service.DBService.GetCompany(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path)
		if err != nil {
			logrus.Panicf("can`t subscribe a topic: %s",err)
		}
		// h.Service.FCM.SendTopicNotification("company", "GET")
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			var data []byte
			if r.Header.Get(("form")) == "expanded" {
				data, err = jsonld.MarshalExpanded(company)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			} else if r.Header.Get(("form")) == "compacted" {
				data, err = jsonld.MarshalCompacted(company)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
			_, err = w.Write(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

	case "PATCH":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPATCH Company")
		logger.Infof("Received request with params %#v", r.URL.Path)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		var company model.Company
		err = jsonld.UnmarshalCompacted(body, &company)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			err = h.Service.DBService.UpdateCompany(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, utils.ToFirestoreMap(company))
			if err != nil {
				h.Service.FCM.SendTopicNotification(path, "Updated")
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

	case "DELETE":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nDELETE Company")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var body map[string][]string
		err := decoder.Decode(&body)
		switch {
		case err == io.EOF:
			err = h.Service.DBService.DeleteCompany(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		case err != nil:
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			if fields, ok := body["fields"]; ok {
				err = h.Service.DBService.DeleteCompanyFields(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, fields)
				if err != nil {
					// TODO render error message with retry option
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}
}

func NewCompanyHandler(svc *service.Service) *CompanyHandler {
	return &CompanyHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*CompanyHandler)(nil)
