package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"path/filepath"
)

var Statuses = []string{"upload"}

type Config struct {
	Env      string `yaml:"env" env-required:"true"`
	Database DBConfig
}

type DBConfig struct {
	DriverName     string `yaml:"driverName" env-required:"true"`
	DataSourceName string `yaml:"dataSourceName" env-required:"true"`
}

func MustLoad() *Config {
	configPath := filepath.Join(os.Getenv("ProgramFiles(x86)"), "JcloudClient", "config", "client.yaml")
	if configPath == "" {
		log.Fatal("Program Files(x86) is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %v", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %v", err)
	}

	return &cfg
}
