package main

import (
	"context"
	"flag"
	"log"
	"os"

	apihttp "github.com/alvarolspeixoto/go-weather-mcp/internal/api/http"
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

	mode := flag.String("mode", "", "Modo de execução: 'http', 'mcp' ou 'dual'")
	port := flag.Int("port", 8080, "Porta do servidor HTTP (apenas para modo http/dual)")
	flag.Parse()

	if *mode == "" {
		*mode = os.Getenv("APP_MODE")
	}
	if *mode == "" {
		*mode = "mcp" // padrão
	}

	// Inicializa dependências
	client := openweathermap.NewClient(
		config.GetApiKey(),
		config.GetWeatherBaseURL(),
		config.GetGeocodeBaseURL(),
	)

	geoRepo := openweathermap.NewGeocodeRepository(client)
	weatherRepo := openweathermap.NewWeatherRepository(client)

	geoUC := geocode.NewUseCase(geoRepo)
	weatherUC := weather.NewUseCase(weatherRepo)

	service := weatherapp.NewService(weatherUC, geoUC)

	switch *mode {
	case "http":
		log.Printf("Iniciando servidor HTTP na porta %d...", *port)
		handler := apihttp.NewHandler(service)
		if err := apihttp.StartServer(handler, *port); err != nil {
			log.Fatalf("Erro ao iniciar servidor HTTP: %v", err)
		}

	case "mcp":
		log.Println("Iniciando MCP Server...")
		weatherTool := tools.NewWeatherTool(service)
		ctx := context.Background()
		if err := mcp.Start(ctx, weatherTool); err != nil {
			log.Fatalf("Erro ao iniciar MCP Server: %v", err)
		}

	case "dual":
		log.Printf("Iniciando MCP e HTTP Server simultaneamente (porta %d)...", *port)

		// Inicia HTTP em goroutine
		go func() {
			handler := apihttp.NewHandler(service)
			if err := apihttp.StartServer(handler, *port); err != nil {
				log.Fatalf("Erro no servidor HTTP: %v", err)
			}
		}()

		// Roda MCP no thread principal
		ctx := context.Background()
		weatherTool := tools.NewWeatherTool(service)
		if err := mcp.Start(ctx, weatherTool); err != nil {
			log.Fatalf("Erro no MCP Server: %v", err)
		}

	default:
		log.Fatalf("Modo inválido: %s. Use 'http', 'mcp' ou 'dual'.", *mode)
	}
}
