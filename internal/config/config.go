package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config - Структура конфигурации приложения
type Config struct {
	Env      string `yaml:"ENV" env-default:"local" env:"ENV"`
	BotToken string `yaml:"BOT_TOKEN" env-required:"true" env:"BOT_TOKEN"`
}

const (
	configPath = "cfg.yaml"
)

// loadConfigFromFile загружает конфигурацию из файла
func loadConfigFromFile(path string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, fmt.Errorf("cannot read config from file: %w", err)
	}
	return &cfg, nil
}

// loadConfigFromEnv загружает конфигурацию из переменных окружения
func loadConfigFromEnv() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("cannot read config from environment: %w", err)
	}
	return &cfg, nil
}

// MustLoad загружает конфигурацию из файла или из переменных окружения
func MustLoad() *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	cfg, err := loadConfigFromEnv()
	if err != nil {
		fmt.Printf("Failed to get environment variables: %v. Trying to load from file %s\n", err, configPath)
		cfg, err = loadConfigFromFile(configPath)
		if err != nil {
			log.Fatalf("cannot load config: %v", err)
		}
	}

	return cfg
}
