package config

import (
	"os"

	"github.com/google/logger"
	"github.com/mcuadros/go-defaults"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	HTTP HTTPConfig
	DB   DB
}

type HTTPConfig struct {
	Address string `default:"0.0.0.0"`
	Port    int    `default:"8080"`
}

type DB struct {
	Type         string `default:"postgres"` // PostgreSQL is only supported DB at the moment
	Address      string `default:"127.0.0.1"`
	Port         int    `default:"5432"`
	User         string
	Password     string
	DatabaseName string
}

func NewConfig(path string) (Config, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		logger.Fatal(err)
	}
	decoder := toml.NewDecoder(file)
	cfg := Config{}
	defaults.SetDefaults(&cfg)
	err = decoder.Decode(&cfg)
	if err != nil {
		logger.Fatal(err)
	}
	return cfg, err
}
