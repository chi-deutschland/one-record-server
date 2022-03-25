package service

type Service struct {
	Env           Env
	SecretManager SecretManager
}

func NewService() *Service {
	svc := Service{}
	return &svc
}
