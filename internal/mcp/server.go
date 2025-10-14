package mcp

import (
	"context"
	"log"

	"github.com/alvarolspeixoto/go-weather-mcp/internal/mcp/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Start(ctx context.Context, weatherTool *tools.WeatherTool) error {
	// Cria o MCP server
	server := mcp.NewServer(&mcp.Implementation{Name: "Weather MCP Server", Version: "v1.0.0"}, nil)

	toolDef := &mcp.Tool{
		Name:        "get_weather_by_city",
		Description: "Retrieves the current weather for a specified city.",
	}

	mcp.AddTool(server, toolDef, weatherTool.GetWeatherByCity)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Error running MCP server: %v", err)
		return err
	}

	return nil
}
