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

type PieceData struct {
	Title   string
	Host    string
	Piece model.Piece
}

type PieceHandler struct {
	Service *service.Service
}

func (h *PieceHandler) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	logger := logrus.WithFields(logrus.Fields{
		"role":       h.Service.Env.ServerRole,
		"request_id": uuid.New().String(),
	})
	logger.Infof("Received request with params %#v", r.URL.Path)

	split_url := strings.Split(r.URL.Path[1:], "/")
	companyID := split_url[0]
	pieceID := split_url[2]
	logger.Debug("Try to fetch piece from DB")
	piece, err := h.Service.DBService.GetPiece(
		h.Service.Env.ProjectId,
		companyID,
		pieceID)
	if err != nil {
		// TODO render error message with retry option
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	logger.Debugf("Fetched piece: %#v", piece)

	json.NewEncoder(w).Encode(piece)
}

// func (h *PieceHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
// 	decoder := json.NewDecoder(r.Body)
// 	var data model.Piece
// 	w.WriteHeader(http.StatusCreated)
// 	logger := logrus.WithFields(logrus.Fields{
// 		"role":       h.Service.Env.ServerRole,
// 		"request_id": uuid.New().String(),
// 	})
// 	logger.Infof("Received request with params %#v", r.URL.Path)

// 	split_url := strings.Split(r.URL.Path[1:], "/")
// 	companyID := split_url[0]
// 	pieceID := split_url[2]
// 	logger.Debug("Try to fetch piece from DB")
// 	piece, err := h.Service.DBService.GetPiece(
// 		h.Service.Env.ProjectId,
// 		companyID,
// 		pieceID)
// 	if err != nil {
// 		// TODO render error message with retry option
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// 	logger.Debugf("Fetched piece: %#v", piece)

// 	json.NewEncoder(w).Encode(piece)
// }

func NewPieceHandler(svc *service.Service) *PieceHandler {
	return &PieceHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*PieceHandler)(nil)
