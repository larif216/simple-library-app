package config

import (
	"net/http"
	"os"
	"time"

	lConfig "simple-library-app/module/library/config"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/subosito/gotenv"
)

type ServiceConfig struct {
	Host string `envconfig:"HOST"`
}

type HttpServer struct {
	HTTPServer *http.Server
	Config     *ServiceConfig
}

func NewHttpServer() (*HttpServer, error) {
	svcCfg, err := loadServiceConfig()
	if err != nil {
		return nil, err
	}

	mux := mux.NewRouter()

	lCfg, err := lConfig.LoadLibraryConfig()
	if err != nil {
		return nil, err
	}

	lCfg.HTTPClient = &http.Client{
		Timeout: 10 * time.Second,
	}

	lConfig.RegisterLibraryHandlers(mux, &lCfg)

	hs := &http.Server{
		Addr:    svcCfg.Host,
		Handler: mux,
	}

	return &HttpServer{
		HTTPServer: hs,
		Config:     &svcCfg,
	}, nil
}

func loadServiceConfig() (ServiceConfig, error) {
	var cfg ServiceConfig

	if _, err := os.Stat(".env"); err == nil {
		if err := gotenv.Load(); err != nil {
			return cfg, err
		}
	}

	err := envconfig.Process("library_service", &cfg)

	return cfg, err
}
