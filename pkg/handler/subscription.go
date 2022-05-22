package handler

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("HERE 0")

	switch r.Method {
	case "POST":
		fmt.Println("HERE 000")

		logger := logrus.WithFields(logrus.Fields{
			"role":       h.Service.Env.ServerRole,
			"request_id": uuid.New().String(),
		})

		fmt.Println("HERE 1")
		logger.Infoln("\nPOST Subscription")
		logger.Infof("Received request with params %#v", r.URL.Path)

		decoder := json.NewDecoder(r.Body)
		fmt.Println("HERE 2")
		var body SubscriptionRequest
		err := decoder.Decode(&body)
		fmt.Println("HERE 3")
		if err != nil {
			// TODO render error message with retry option
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			fmt.Println("HERE 5")
			fmt.Printf("\n%#v", body)
			res, err := h.Service.FCM.Subscribe(body.Topic, []string{body.Token})
			fmt.Printf("\n-->%#v", res)
			fmt.Println("HERE 6")
			fmt.Printf("\n%#v\n", err)
			if err != nil {
				// TODO render error message with retry option
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				json.NewEncoder(w).Encode(res)
			}
		}
	}
}

func NewSubscriptionHandler(svc *service.Service) *SubscriptionHandler {
	return &SubscriptionHandler{Service: svc}
}

var _ onerecordhttp.ContextHandler = (*SubscriptionHandler)(nil)
