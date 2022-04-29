package handler

// import (
// 	"net/http"
// 	"strings"
// 	"github.com/chi-deutschland/one-record-server/pkg/model"
// 	"github.com/chi-deutschland/one-record-server/pkg/service"
// 	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
// 	"github.com/google/uuid"
// 	"github.com/sirupsen/logrus"
//     "encoding/json"
// )

// type ShipmentData struct {
// 	Title   string
// 	Host    string
// 	Shipment model.Shipment
// }

// type ShipmentHandler struct {
// 	Service *service.Service
// }

// func (h *ShipmentHandler) Handler(w http.ResponseWriter, r *http.Request) {
// 	logger := logrus.WithFields(logrus.Fields{
// 		"role":       h.Service.Env.ServerRole,
// 		"request_id": uuid.New().String(),
// 	})
// 	logger.Infof("Received request with params %#v", r.URL.Path)
	
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)

// 	split_url := strings.Split(r.URL.Path[1:], "/")
// 	companyID := split_url[0]
// 	pieceID := split_url[2]
// 	logger.Debug("Try to fetch shipment from DB")
// 	shipment, err := h.Service.DBService.GetShipment(
// 		h.Service.Env.ProjectId,
// 		companyID,
// 		pieceID)
// 	if err != nil {
// 		// TODO render error message with retry option
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// 	logger.Debugf("Fetched shipment: %#v", shipment)

// 	json.NewEncoder(w).Encode(shipment)
// }

// func NewShipmentHandler(svc *service.Service) *ShipmentHandler {
// 	return &ShipmentHandler{Service: svc}
// }

// var _ onerecordhttp.ContextHandler = (*ShipmentHandler)(nil)
