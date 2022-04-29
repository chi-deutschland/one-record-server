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
	w.Header().Set("Content-Type", "text/html")
	logger := logrus.WithFields(logrus.Fields{
		"role":       h.Service.Env.ServerRole,
		"request_id": uuid.New().String(),
	})
	logger.Infof("Received request with params %#v", r.URL.Path)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	companyID := strings.Split(r.URL.Path[1:], "/")[0]
	pieces, err := h.Service.DBService.GetPieces(
		h.Service.Env.ProjectId,
		companyID)
	if err != nil {
		// TODO render error message with retry option
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	logger.Debugf("Fetched pieces: %#v", pieces)

	json.NewEncoder(w).Encode(pieces)
}

func NewPiecesHandler(svc *service.Service) *PiecesHandler {
	return &PiecesHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*PiecesHandler)(nil)
