package gcp

import (
	"context"
	"fmt"
	"sync/atomic"

	"cloud.google.com/go/pubsub"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
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



	// fmt.Println("Listing all topics from the project:")


	subID := "my-sub"
	
	go ps.PullMsgs(subID)
	return &ps, nil
}

func list(client *pubsub.Client) ([]*pubsub.Topic, error) {
	ctx := context.Background()
	var topics []*pubsub.Topic
	it := client.Topics(ctx)
	for {
		topic, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}
	return topics, nil
}

func (ps *PS) Publish(topic string, message string) error {

        
	t := ps.Client.Topic(topic)
	fmt.Println(1,topic,message,ps.Ctx,t)
	result := t.Publish(ps.Ctx, &pubsub.Message{
                Data: []byte(message),
	})
        fmt.Println(2,topic,message,ps.Ctx,t,result)
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ps.Ctx)
        fmt.Println(3,topic,message,ps.Ctx,t,result,id,err)
	if err != nil {
                panic(err)
		// return fmt.Errorf("pubsub: result.Get: %v", err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)


	// t := ps.Client.Topic(topic)
        // result := t.Publish(ps.Ctx, &pubsub.Message{
        //         Data: []byte(message),
        // })
        // // Block until the result is returned and a server-generated
        // // ID is returned for the published message.
        // id, err := result.Get(ps.Ctx)
        // if err != nil {
        //         return fmt.Errorf("pubsub: result.Get: %v", err)
        // }
        // fmt.Printf( "Published a message; msg ID: %v\n", id)
        return nil
	
}

func (ps *PS) PullMsgs(subID string) error {

        fmt.Println("PullMsgs")
        var received int32
        err := ps.Client.Subscription(subID).Receive(ps.Ctx, func(_ context.Context, msg *pubsub.Message) {
                fmt.Printf("\n\n-->Got message: %q\n\n", string(msg.Data))
                atomic.AddInt32(&received, 1)
                msg.Ack()
        })
        if err != nil {
                return fmt.Errorf("sub.Receive: %v", err)
        }
        fmt.Printf("Received %d messages\n", received)

        return nil
}