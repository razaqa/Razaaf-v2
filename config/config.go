package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var appConfig AppConfig

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("error loading .env file")
		os.Exit(2)
	}

	appConfig = AppConfig{
		Port:			os.Getenv("PORT"),
		Environment:	os.Getenv("ENVIRONMENT"),
	}
}

func GetAppConfig() AppConfig {
	return appConfig
}