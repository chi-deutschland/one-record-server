
package gcp

import (
	"context"

	"github.com/sirupsen/logrus"
	"cloud.google.com/go/pubsub"
)

type PS struct {
	Client *pubsub.Client
	Ctx context.Context
}

func NewPubSub() (*PS, error) {
	var ps PS
	var err error
	ps.Ctx = context.Background()
	projectID := "one-record"
	ps.Client, err = pubsub.NewClient(ps.Ctx, projectID)
	if err != nil {
		logrus.Panicf("can`t create client pub/sub: %s",err)
	}
	defer ps.Client.Close()
	return &ps, nil
}


func (ps *PS) publish(topic string, message string)  {
	
}