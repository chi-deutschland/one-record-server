package service

type Service struct {
	Env           Env
	SecretManager SecretManager
	DBService     DBService
}

func NewService() *Service {
	svc := Service{}
	return &svc
}
