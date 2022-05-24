package main

import (
	"github.com/alseiitov/yacb/service_crypto_currency/config"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/app"
	"log"
	"os"
)

func main() {

	b, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.ParseConfig(b)
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
