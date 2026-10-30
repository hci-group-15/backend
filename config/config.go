package config

type ConfigUrls struct {
	FrontendURL     string
	BackendURL      string
	DefaultRedirect string
}

type Config struct {
	Urls ConfigUrls
}

func Load() Config {
	return Config{
		Urls: ConfigUrls{
			FrontendURL: "http://localhost:8081",
			DefaultRedirect: "http://localhost:8081",
			BackendURL: "http://localhost:8080",
		},
	}
}
