package utils

import "github.com/chi-deutschland/one-record-server/pkg/service"

type AuthHeaderSecretValues struct {
	Key, Value string
}

func NewAuthHeaderSecretValues(svc *service.Service) (*AuthHeaderSecretValues, error) {
	var h AuthHeaderSecretValues
	ver := "latest"
	k, err := svc.SecretManager.GetSecret(svc.Env.Auth.Key, ver, svc.Env.ProjectId)
	if err != nil {
		return &h, err
	}
	h.Key = string(k)
	v, err := svc.SecretManager.GetSecret(svc.Env.Auth.Value, ver, svc.Env.ProjectId)
	if err != nil {
		return &h, err
	}
	h.Value = string(v)
	return &h, nil
}
