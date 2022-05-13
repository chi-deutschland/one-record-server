package handler

import (
	"encoding/json"
	"net/http"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type SubscriptionHandler struct {
	Service *service.Service
}

type SubscriptionRequest struct {
	Topic string `firestore:"topic,omitempty" json:"topic,omitempty"`
	Token string `firestore:"token,omitempty" json:"token,omitempty"`
}

func (h *SubscriptionHandler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})
		logger.Infoln("\nPOST Subscription")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		var body SubscriptionRequest
		err := decoder.Decode(&body)
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			_, err := h.Service.FCM.Subscribe(body.Topic, []string{body.Token})
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

func NewSubscriptionHandler(svc *service.Service) *SubscriptionHandler {
	return &SubscriptionHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*SubscriptionHandler)(nil)