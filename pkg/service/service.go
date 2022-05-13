package service

type Service struct {
	Env           Env
	SecretManager SecretManager
	DBService     DBService
	FCM			  FCM
}

func NewService() *Service {
	svc := Service{}
	return &svc
}
