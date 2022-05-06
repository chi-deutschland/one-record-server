package handler

import (
	"encoding/json"
	"net/http"

	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type PiecesHandler struct {
	Service *service.Service
}

func (h *PiecesHandler) Handler(w http.ResponseWriter, r *http.Request) {
	path := PathMultipleEntries(r.URL.Path)
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json+ld")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nGET Pieces")
		logger.Infof("Received request with params %#v", r.URL.Path)

		pieces, err := h.Service.DBService.GetPieces(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			json.NewEncoder(w).Encode(pieces)
		}

	case "POST":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPOST Piece")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var piece model.Piece
		err := decoder.Decode(&piece)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			ID, err := h.Service.DBService.AddPiece(h.Service.Env.ProjectId, h.Service.Env.ServerRole, path, piece.ID, piece)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				json.NewEncoder(w).Encode(map[string]string{"id": ID})
				w.WriteHeader(http.StatusCreated)
			}
		}
	}
}

func NewPiecesHandler(svc *service.Service) *PiecesHandler {
	return &PiecesHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*PiecesHandler)(nil)
