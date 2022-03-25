package service

type Env struct {
	Auth struct {
		Value string `env:"AUTH_VALUE"`
		Key   string `env:"AUTH_KEY"`
	}
	Path struct {
		Template string `env:"PATH_TEMPLATE"`
		Static   string `env:"PATH_STATIC"`
	}
	ServerRole string `env:"SRV_ROLE"`
	ProjectId  string `env:"PROJECT_ID"`
}
