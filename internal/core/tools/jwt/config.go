package tools_jwt

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Secret []byte `envconfig:"SECRET" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("JWT", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}

	if len(config.Secret) == 0 {
		return Config{}, fmt.Errorf("JWT secret cannot be empty")
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get jwt config: %w", err)
		panic(err)
	}
	return config
}
