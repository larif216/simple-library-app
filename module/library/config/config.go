package config

import (
	"net/http"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/subosito/gotenv"
)

type LibraryConfig struct {
	BaseURL    string       `envconfig:"BOOK_URL"`
	HTTPClient *http.Client `ignore:"true"`
}

func LoadLibraryConfig() (LibraryConfig, error) {
	var cfg LibraryConfig

	if _, err := os.Stat(".env"); err == nil {
		if err := gotenv.Load(); err != nil {
			return cfg, err
		}
	}

	err := envconfig.Process("library_service", &cfg)

	return cfg, err
}
