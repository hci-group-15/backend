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
	return Config{}
}
