package service

type Service struct {
	Env           Env
	SecretManager SecretManager
	DBService     DBService
	FCM			  FCM
	PS			  PS
}

func NewService() *Service {
	svc := Service{}
	return &svc
}
