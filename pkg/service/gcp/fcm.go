package gcp

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"

	"fmt"
)

type FCM struct {
	Client *messaging.Client
	Ctx    context.Context
}

func NewFCM() (*FCM, error) {
	var f FCM
	var err error
	f.Ctx, f.Client, err = getClient()
	fmt.Printf("\n\n%#v\n\n%#v\n\n", f, err)
	if err != nil {
		return &f, err
	}

	return &f, nil
}

var _ service.FCM = (*FCM)(nil)

func (f *FCM) SendTopicNotification(topic string, status string) (response string, err error) {
	notification := &messaging.Notification{
		Title: topic + " " + status,
		Body:  "Body",
	}

	message := &messaging.Message{
		Data: map[string]string{
			"data": "value",
		},
		Notification: notification,
		Topic:        topic,
	}

	fmt.Println("\n*************** message", message, topic)

	// Send a message to the devices subscribed to the provided topic.
	response, err = f.Client.Send(f.Ctx, message)
	if err != nil {
		logrus.Panicf("can`t send notification: %s", err)
	}
	return response, err
}

func (f *FCM) Subscribe(topic string, tokens []string) (response *messaging.TopicManagementResponse, err error) {
	response, err = f.Client.SubscribeToTopic(f.Ctx, tokens, topic)
	if err != nil {
		logrus.Panicf("can`t subscribe a topic: %s", err)
	}
	return response, err

}

func getClient() (ctx context.Context, client *messaging.Client, err error) {
	// config := &firebase.Config{ProjectID: "one-record"}

	// app, err := firebase.NewApp(context.Background(), config)
	// if err != nil {
	// 	return ctx, client, err
	// }

	opt := option.WithCredentialsFile("../one-record-firebase-adminsdk-l8dmp-98c472bf32.json")
	// opt := option.with
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logrus.Panicf("can`t subscribe a topic: %s", err)
	}

	client, err = app.Messaging(context.Background())
	if err != nil {
		return ctx, client, err
	}

	return context.Background(), client, nil
}
