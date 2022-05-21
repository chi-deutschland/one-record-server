package service

type PS interface {
	Publish(topic string, message string) error
	PullMsgs(subID string)error
}
