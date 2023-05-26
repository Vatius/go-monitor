package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App
		HTTP
		Log
		PG
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		//PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		Host     string `env-required:"true"                 env:"PG_HOST"`
		Port     string `env-required:"true"                 env:"PG_PORT"`
		User     string `env-required:"true"                 env:"PG_USER"`
		Password string `env-required:"true"                 env:"PG_PASSWORD"`
		DbName   string `env-required:"true"                 env:"PG_DBNAME"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
