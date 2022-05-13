
package gcp

import (
	"github.com/chi-deutschland/one-record-server/pkg/service"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
    "google.golang.org/api/option"
	"context"
)

type FCM struct {
	Client *messaging.Client
	Ctx context.Context
}

func NewFCM() (*FCM, error) {
	var f FCM
	var err error
	f.Ctx, f.Client, err = GetClient()
	
	if err != nil {
		return &f, err
	}

	return &f, nil
}

var _ service.FCM = (*FCM)(nil)

func (f *FCM) SendTopicNotification(topic string, status string) (response string, err error) {
	notification := &messaging.Notification{
        Title: topic + " " + status,
        Body: "Body",
    }

    message := &messaging.Message{
        Data: map[string]string{
            "data": "value",
        },
        Notification: notification,
        Topic: topic,
    }

    // Send a message to the devices subscribed to the provided topic.
    response, err = f.Client.Send(f.Ctx, message)

    return response, err
}

func (f *FCM) Subscribe(topic string, tokens []string) (response *messaging.TopicManagementResponse, err error) {
    response, err = f.Client.SubscribeToTopic(f.Ctx, tokens, topic)
    return response, err
}

func GetClient() (ctx context.Context, client *messaging.Client, err error) {
	app, err := firebase.NewApp(context.Background(), nil, option.WithoutAuthentication())
	if err != nil {
		return ctx, client, err
	}

	client, err = app.Messaging(context.Background())
	if err != nil {
		return ctx, client, err
	}

	return ctx, client, nil
}