package service

type SecretManager interface {
	GetSecret(
		secret,
		version,
		projectID string,
	) (
		[]byte,
		error,
	)
}
