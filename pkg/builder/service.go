package builder

import (
	"github.com/chi-deutschland/one-record-server/pkg/service"
)

type ServiceBuilder struct {
	service *service.Service
}

func NewServiceBuilder() *ServiceBuilder {
	return &ServiceBuilder{service: &service.Service{}}
}

func (s *ServiceBuilder) WithEnv(env service.Env) *ServiceBuilder {
	s.service.Env = env
	return s
}

func (s *ServiceBuilder) WithGcpSecretManager(secretManager service.SecretManager) *ServiceBuilder {
	s.service.SecretManager = secretManager
	return s
}

func (s *ServiceBuilder) WithGcpFirestore(firestore service.DBService) *ServiceBuilder {
	s.service.DBService = firestore
	return s
}

func (s *ServiceBuilder) Build() *service.Service {
	return s.service
}
