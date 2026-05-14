package core_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type HTTPServerConfig struct {
	Address          string        `envconfig:"ADDRESS" required:"true"`
	ShutdownDuration time.Duration `envconfig:"SHUTDOWN_DURATION" default:"30s"`
	StaticURL        string        `envconfig:"STATIC_DIR" default:"/static"`
}

func NewHTTPServerConfig() (*HTTPServerConfig, error) {
	var config HTTPServerConfig
	if err := envconfig.Process("HTTP", &config); err != nil {
		return &HTTPServerConfig{}, fmt.Errorf("envconfig process: %w", err)
	}
	return &config, nil
}

func NewHTTPServerConfigMust() *HTTPServerConfig {
	config, err := NewHTTPServerConfig()
	if err != nil {
		err = fmt.Errorf("get http server config: %w", err)
		panic(err)
	}
	return config
}
