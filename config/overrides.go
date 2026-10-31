package config

import "os"

func loadEnvVarSettings(conf *Config) {
	frontendUrl, ok := os.LookupEnv("FRONTEND_URL")
	if ok {
		(*conf).Urls.FrontendURL = frontendUrl
	}
	backendUrl, ok := os.LookupEnv("BACKEND_URL")
	if ok {
		(*conf).Urls.BackendURL = backendUrl
	}
	defaultRedirect, ok := os.LookupEnv("DEFAULT_REDIRECT")
	if ok {
		(*conf).Urls.DefaultRedirect = defaultRedirect
	}
	dataDir, ok := os.LookupEnv("DATA_DIR")
	if ok {
		(*conf).Datadir = dataDir
	}
}
