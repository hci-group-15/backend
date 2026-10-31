package config

import "github.com/TwiN/logr"

func checkConfig(conf *Config) {
	if (*conf).Datadir == "" {
		(*conf).Datadir = "./data/"
	}
	if (*conf).Datadir[len((*conf).Datadir)-1] != '/' {
		(*conf).Datadir += "/"
	}

	if (*conf).Urls.BackendURL == "" {
		logr.Warn("Backend URL was not configured, falling back to localhost")
		(*conf).Urls.BackendURL = "http://localhost:8080"
	}
	if (*conf).Urls.FrontendURL == "" {
		logr.Warn("Frontend URL URL was not configured, falling back to localhost")
		(*conf).Urls.FrontendURL = "http://localhost:8081"
	}
	if (*conf).Urls.DefaultRedirect == "" {
		logr.Warn("DefaultRedirect was not configured, falling back to default")
		(*conf).Urls.DefaultRedirect = conf.Urls.FrontendURL + "/app"
	}

	logr.Info("Configuration loaded and validated successfully")
}
