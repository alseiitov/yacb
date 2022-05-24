package main

import (
	"github.com/alseiitov/yacb/service_telegram_bot/config"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/app"
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
