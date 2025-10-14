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

func GetWeatherBaseURL() string {
	url := os.Getenv("OPENWEATHER_WEATHER_URL")
	if url == "" {
		url = "https://api.openweathermap.org/data/2.5/weather"
	}
	return url
}

func GetGeocodeBaseURL() string {
	url := os.Getenv("OPENWEATHER_GEO_URL")
	if url == "" {
		url = "http://api.openweathermap.org/geo/1.0/direct"
	}
	return url
}
