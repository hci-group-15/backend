package config

import (
	"os"

	"github.com/TwiN/logr"
	"github.com/goccy/go-yaml"
)

type ApplicationUrls struct {
	FrontendURL     string   `json:"frontendURL"`
	DefaultRedirect string   `json:"defaultRedirect"`
	BackendURL      string   `json:"backendURL"`
	TrustedProxies  []string `json:"trustedProxies"`
}

type Config struct {
	Datadir     string          `json:"dataDir"`
	Urls        ApplicationUrls `json:"urls"`
	TestingMode bool
}

func Load() Config {
	var conf Config

	// Load config
	contents, err := os.ReadDir("./")
	fname := "config.yml"
	for _, file := range contents {
		if file.Name() == "config.secret.yml" {
			fname = "config.secret.yml"
		}
	}
	data, err := os.ReadFile(fname)
	if err != nil {
		logr.Fatal("[FATAL] Failed to load configuration")
	}
	yaml.Unmarshal(data, &conf)

	// Load overrides form env vars
	loadEnvVarSettings(&conf)

	// Check config
	checkConfig(&conf)

	return conf
}
