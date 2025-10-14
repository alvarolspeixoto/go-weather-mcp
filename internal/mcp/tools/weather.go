package tools

import (
	"context"
	"fmt"

	"github.com/alvarolspeixoto/go-weather-mcp/internal/app/weather"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type WeatherTool struct {
	svc *weather.Service
}

func NewWeatherTool(svc *weather.Service) *WeatherTool {
	return &WeatherTool{
		svc: svc,
	}
}

func (t *WeatherTool) GetWeatherByCity(ctx context.Context, req *mcp.CallToolRequest, input map[string]any) (*mcp.CallToolResult, any, error) {
	city, ok := input["city"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'city'"), nil
	}

	data, err := t.svc.GetWeatherByCityName(ctx, city)
	if err != nil {
		return nil, err, nil
	}

	return nil, map[string]any{
		"city":        data.CityName,
		"country":     data.Country,
		"temperature": data.Temperature,
		"description": data.Description,
	}, nil
}
