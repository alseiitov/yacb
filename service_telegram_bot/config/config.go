package config

import (
	"github.com/alseiitov/yacb/pkg/postgres"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	CryptoCurrencyService CryptoCurrencyService `yaml:"crypto-currency-service"`
	Telegram              Telegram              `yaml:"telegram"`
	Postgres              postgres.Config       `yaml:"postgres"`
}

type CryptoCurrencyService struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Telegram struct {
	APIKey        string `yaml:"api-key"`
	UpdateTimeout int    `yaml:"update-timeout"`
}

func ParseConfig(fileBytes []byte) (*Config, error) {

	var cfg Config

	b := []byte(os.ExpandEnv(string(fileBytes)))
	err := yaml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
