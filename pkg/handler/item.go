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

type ItemHandler struct {
	Service *service.Service
}

func (h *ItemHandler) Handler(w http.ResponseWriter, r *http.Request) {
	path := PathMultipleEntries(r.URL.Path)
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nGET Item")
		logger.Infof("Received request with params %#v", r.URL.Path)

		item, err := h.Service.DBService.GetItem(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			var data []byte
			if r.Header.Get(("form")) == "expanded" {
				data, err = jsonld.MarshalExpanded(item)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			} else if r.Header.Get(("form")) == "compacted" {
				data, err = jsonld.MarshalCompacted(item)
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
		logger.Infoln("\nPATCH Item")
		logger.Infof("Received request with params %#v", r.URL.Path)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		var item model.Item
		err = jsonld.UnmarshalCompacted(body, &item)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			err = h.Service.DBService.UpdateItem(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, utils.ToFirestoreMap(item))
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

	case "DELETE":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nDELETE Item")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var body map[string][]string
		err := decoder.Decode(&body)
		switch {
		case err == io.EOF:
			err = h.Service.DBService.DeleteItem(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		case err != nil:
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			if fields, ok := body["fields"]; ok {
				err = h.Service.DBService.DeleteItemFields(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, fields)
				if err != nil {
					// TODO render error message with retry option
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}
}

func NewItemHandler(svc *service.Service) *ItemHandler {
	return &ItemHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*ItemHandler)(nil)