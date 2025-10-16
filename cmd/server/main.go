package main

import (
	"context"
	"log"

	weatherapp "github.com/alvarolspeixoto/go-weather-mcp/internal/app/weather"
	"github.com/alvarolspeixoto/go-weather-mcp/internal/config"
	"github.com/alvarolspeixoto/go-weather-mcp/internal/domain/geocode"
	"github.com/alvarolspeixoto/go-weather-mcp/internal/domain/weather"
	"github.com/alvarolspeixoto/go-weather-mcp/internal/infra/openweathermap"
	"github.com/alvarolspeixoto/go-weather-mcp/internal/mcp"
	"github.com/alvarolspeixoto/go-weather-mcp/internal/mcp/tools"
)

func main() {
	config.Load()

	// Cria o client OpenWeather
	client := openweathermap.NewClient(config.GetApiKey(), config.GetWeatherBaseURL(), config.GetGeocodeBaseURL())

	geoRepo := openweathermap.NewGeocodeRepository(client)
	weatherRepo := openweathermap.NewWeatherRepository(client)

	geoUC := geocode.NewUseCase(geoRepo)
	weatherUC := weather.NewUseCase(weatherRepo)

	service := weatherapp.NewService(weatherUC, geoUC)
	weatherTool := tools.NewWeatherTool(service)

	ctx := context.Background()

	log.Println("Starting MCP Server...")

	if err := mcp.Start(ctx, weatherTool); err != nil {
		log.Fatal(err)
	}
}
