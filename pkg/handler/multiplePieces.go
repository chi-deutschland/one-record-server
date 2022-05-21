package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/chi-deutschland/one-record-server/pkg/jsonld"
	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type MultiplePiecesHandler struct {
	Service *service.Service
}

func (h *MultiplePiecesHandler) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json+ld")
	logger := logrus.WithFields(logrus.Fields{
		"role":       h.Service.Env.ServerRole,
		"request_id": uuid.New().String(),
	})
	logger.Debugln("\nGET Multiple Pieces")
	logger.Infof("Received request with params %#v", r.URL.Path)

	decoder := json.NewDecoder(r.Body)
	var body map[string][]string
	var pieces []model.Piece
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if paths, ok := body["ids"]; ok {
		for _, path := range paths {
			// piece, err := h.Service.DBService.GetPiece(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path)

			client := &http.Client{
				Timeout: time.Second * 10,
			}
			req, err := http.NewRequest("GET", path, nil)
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
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			pieces = append(pieces, piece)
		}
		data, err := jsonld.MarshalCompacted(pieces)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		_, err = w.Write(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func NewMultiplePiecesHandler(svc *service.Service) *MultiplePiecesHandler {
	return &MultiplePiecesHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*MultiplePiecesHandler)(nil)
