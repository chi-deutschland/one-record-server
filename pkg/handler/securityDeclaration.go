package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Meschkov/jsonld"
	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	utils "github.com/chi-deutschland/one-record-server/pkg/utils/conv"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type SecurityDeclarationHandler struct {
	Service *service.Service
}

func (h *SecurityDeclarationHandler) Handler(w http.ResponseWriter, r *http.Request) {
	colPath, docPath, id := PathSingleEntry(r.URL.Path, "securityDeclarations")
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nGET SecurityDeclaration")
		logger.Infof("Received request with params %#v", r.URL.Path)

		securityDeclaration, err := h.Service.DBService.GetSecurityDeclaration(h.Service.Env.ProjectId, h.Service.Env.ServerRole, docPath)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			var data []byte
			if r.Header.Get(("form")) == "expanded" {
				data, err = jsonld.MarshalExpanded(securityDeclaration)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			} else if r.Header.Get(("form")) == "compacted" {
				data, err = jsonld.MarshalCompacted(securityDeclaration)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
			_, err = w.Write(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

	case "POST":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPOST SecurityDeclaration")
		logger.Infof("Received request with params %#v", r.URL.Path)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		var objmap map[string]json.RawMessage
		err = json.Unmarshal(body, &objmap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		var pieceId string
		err = json.Unmarshal(objmap["pieceId"], &pieceId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		var securityDeclaration model.SecurityDeclaration
		err = jsonld.UnmarshalCompacted(objmap["obj"], &securityDeclaration)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {

			ID, err := h.Service.DBService.AddSecurityDeclaration(h.Service.Env.ProjectId, h.Service.Env.ServerRole, colPath, id, securityDeclaration)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {

				json.NewEncoder(w).Encode(map[string]string{"id": ID})

				// PATCH to Piece Server

				client := &http.Client{
					Timeout: time.Second * 10,
				}

				data := url.Values{}
				data.Set("https://onerecord.iata.org/cargo#piece#securityDeclaration", r.URL.RequestURI())

				l := len(r.URL.Path) - 21

				s := r.URL.Path[:l]
				fmt.Println("http://localhost:8081/" + s)
				req, err := http.NewRequest("PATCH", "http://localhost:8081/"+s, strings.NewReader(data.Encode()))
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				req.Header.Set("form", "compacted")
				req.Header.Set("x-auth-name", "dr6YEPzk6zPyLG2GXvwF3ZcdBYUyTRq8MwHM8hJBfWY9sXCiGz")
				response, err := client.Do(req)

				// decoder := json.NewDecoder(response.Body)
				var piece model.Piece
				body, err := ioutil.ReadAll(response.Body)

				jsonld.UnmarshalCompacted(body, &piece)
				// Pub/Sub

				w.WriteHeader(http.StatusCreated)
			}
		}

	case "PATCH":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPATCH SecurityDeclaration")
		logger.Infof("Received request with params %#v", r.URL.Path)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		var securityDeclaration model.SecurityDeclaration
		err = jsonld.UnmarshalCompacted(body, &securityDeclaration)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			err = h.Service.DBService.UpdateSecurityDeclaration(h.Service.Env.ProjectId, h.Service.Env.ServerRole, docPath, utils.ToFirestoreMap(securityDeclaration))
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
		logger.Infoln("\nDELETE SecurityDeclaration")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var body map[string][]string
		err := decoder.Decode(&body)
		switch {
		case err == io.EOF:
			err = h.Service.DBService.DeleteSecurityDeclaration(h.Service.Env.ProjectId, h.Service.Env.ServerRole, docPath)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		case err != nil:
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			if fields, ok := body["fields"]; ok {
				err = h.Service.DBService.DeleteSecurityDeclarationFields(h.Service.Env.ProjectId, h.Service.Env.ServerRole, docPath, fields)
				if err != nil {
					// TODO render error message with retry option
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}
}

func NewSecurityDeclarationHandler(svc *service.Service) *SecurityDeclarationHandler {
	return &SecurityDeclarationHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*SecurityDeclarationHandler)(nil)
