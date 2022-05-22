package service

import (
	"firebase.google.com/go/messaging"
)

type FCM interface {
	SendTopicNotification(topic string, status string) (response string, err error)
	Subscribe(topic string, tokens []string) (response *messaging.TopicManagementResponse, err error)
}
