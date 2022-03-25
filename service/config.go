package service

type Service struct {
	Auth struct {
		NameValue  string `env:"AUTH_NAME_VALUE"`
		NameKey    string `env:"AUTH_NAME_KEY"`
		TokenKey   string `env:"AUTH_TOKEN_KEY"`
		TokenValue string `env:"AUTH_TOKEN_VALUE"`
	}
	Path struct {
		Template string `env:"PATH_TEMPLATE"`
		Static   string `env:"PATH_STATIC"`
	}
	ServerRole string `env:"SRV_ROLE"`
	ProjectId  string `env:"PROJECT_ID"`
}
