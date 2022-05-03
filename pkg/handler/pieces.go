package handler

import (
	"net/http"
	"strings"
	"github.com/chi-deutschland/one-record-server/pkg/model"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
    "encoding/json"
)

type PiecesData struct {
	Title   string
	Host    string
	Pieces []model.Piece
}

type PiecesHandler struct {
	Service *service.Service
}

func (h *PiecesHandler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json+ld")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})

		logger.Debugln("\nGET PIECES")
		companyID := strings.Split(r.URL.Path[1:], "/")[0]
		pieces, err := h.Service.DBService.GetPieces(
		h.Service.Env.ProjectId,
		companyID)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			logger.Debugf("Fetched pieces: %#v", pieces)
			json.NewEncoder(w).Encode(pieces)
		}

	case "POST":
		w.Header().Set("Content-Type", "application/json+ld")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Debugln("\nPOST PIECE")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var piece model.Piece
		err := decoder.Decode(&piece)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			companyID := strings.Split(r.URL.Path[1:], "/")[0]
			var pieceID, err = h.Service.DBService.AddPiece(
			h.Service.Env.ProjectId, companyID, piece)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				json.NewEncoder(w).Encode(map[string]string{"id": pieceID})
				w.WriteHeader(http.StatusCreated)
			}
		}
	}
}

func NewPiecesHandler(svc *service.Service) *PiecesHandler {
	return &PiecesHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*PiecesHandler)(nil)
