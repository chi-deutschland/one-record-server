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
	"io"
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
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nGET PIECE")
		logger.Infof("Received request with params %#v", r.URL.Path)

		split_url := strings.Split(r.URL.Path[1:], "/")
		companyID := split_url[0]
		pieceID := split_url[2]
		piece, err := h.Service.DBService.GetPiece(
			h.Service.Env.ProjectId,
			companyID,
			pieceID)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			logger.Debugf("Fetched piece: %#v", piece)
			json.NewEncoder(w).Encode(piece)
		}

	case "PATCH":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPATCH PIECE")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var piece model.Piece
		err := decoder.Decode(&piece)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			split_url := strings.Split(r.URL.Path[1:], "/")
			companyID := split_url[0]
			pieceID := split_url[2]
			logger.Debugln(companyID, pieceID)
			logger.Debugf("piece: %#v", piece)
			h.Service.DBService.UpdatePiece(
			h.Service.Env.ProjectId,
			companyID, pieceID, piece)
		}

	case "DELETE":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nDELETE PIECE")
		logger.Infof("Received request with params %#v", r.URL.Path)

		split_url := strings.Split(r.URL.Path[1:], "/")
		companyID := split_url[0]
		pieceID := split_url[2]
		decoder := json.NewDecoder(r.Body)
		var body map[string][]string
		err := decoder.Decode(&body)
		switch {
		case err == io.EOF:
			h.Service.DBService.DeletePiece(
				h.Service.Env.ProjectId,
				companyID, pieceID)
		case err != nil:
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			if fields, ok := body["fields"]; ok {
				h.Service.DBService.DeletePieceFields(
				h.Service.Env.ProjectId,
				companyID, pieceID, fields)
			}
		}
	}
}

func NewPieceHandler(svc *service.Service) *PieceHandler {
	return &PieceHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*PieceHandler)(nil)
