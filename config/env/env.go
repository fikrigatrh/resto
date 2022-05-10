package env

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"resto/models"
)

var Config models.ServerConfig

func init() {
	err := loadConfig()
	if err != nil {
		log.Fatalf("cannot open file .env.local")
	}
}

func loadConfig() error {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err, "config/env: load file env")
	}

	err = env.Parse(&Config)
	if err != nil {
		return err
	}

	return err
}
