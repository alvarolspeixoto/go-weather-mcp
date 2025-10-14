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

	// Cria os repositórios
	geoRepo := openweathermap.NewGeocodeRepository(client)
	weatherRepo := openweathermap.NewWeatherRepository(client)

	// Cria usecases
	geoUC := geocode.NewUseCase(geoRepo)
	weatherUC := weather.NewUseCase(weatherRepo)

	// Cria service de aplicação
	service := weatherapp.NewService(weatherUC, geoUC)

	// Cria a tool
	weatherTool := tools.NewWeatherTool(service)

	// Cria context
	ctx := context.Background()

	log.Println("Iniciando o servidor MCP...")

	// Inicia o MCP server
	if err := mcp.Start(ctx, weatherTool); err != nil {
		log.Fatal(err)
	}

	log.Println("Servidor MCP rodando com sucesso.")
}
