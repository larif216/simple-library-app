package config

import "net/http"

type LibraryConfig struct {
	BaseURL    string
	HTTPClient *http.Client
}

func LoadLibraryConfig() *LibraryConfig {
	return new(LibraryConfig)
}
