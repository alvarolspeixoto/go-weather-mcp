package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, proceeding with system environment variables")
	}
}

func GetApiKey() string {
	return os.Getenv("OPENWEATHER_API_KEY")
}