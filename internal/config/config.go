package config

import (
	"os"

	"github.com/kkyr/fig"
)

type Config struct {
	Server Server `fig:"server" required:"true"`
}

type Server struct {
	Port int `fig:"port" default:"8080"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := fig.Load(
		&cfg,
		fig.File(LookUpEnvOrDefault("CONFIG_NAME", "config.yaml")),
		fig.Dirs(LookUpEnvOrDefault("CONFIG_BASE_DIR", "/opt/jchapman63/gkmn")),
		fig.UseEnv(LookUpEnvOrDefault("ENV_PREFIX", "api")),
	); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func LookUpEnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
