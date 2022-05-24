package config

import (
	"github.com/alseiitov/yacb/pkg/postgres"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server   Server          `yaml:"server"`
	Binance  Binance         `yaml:"binance"`
	Postgres postgres.Config `yaml:"postgres"`
}

type Server struct {
	Port        int `yaml:"port"`
	GatewayPort int `yaml:"gateway-port"`
}

type Binance struct {
	URL string `yaml:"url"`
}

func ParseConfig(fileBytes []byte) (*Config, error) {

	var cf Config

	b := []byte(os.ExpandEnv(string(fileBytes)))
	err := yaml.Unmarshal(b, &cf)
	if err != nil {
		return nil, err
	}

	return &cf, nil
}
