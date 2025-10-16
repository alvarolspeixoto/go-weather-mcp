package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/alvarolspeixoto/go-weather-mcp/internal/app/weather"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type WeatherTool struct {
	svc *weather.Service
}

func NewWeatherTool(svc *weather.Service) *WeatherTool {
	return &WeatherTool{svc: svc}
}

func (t *WeatherTool) GetWeatherByCity(ctx context.Context, req *mcp.CallToolRequest, input map[string]any) (*mcp.CallToolResult, any, error) {
	city, ok := input["city"].(string)
	if !ok || city == "" {
		return nil, nil, fmt.Errorf("missing or invalid 'city' field")
	}

	log.Printf("[MCP] Recebida requisição de clima para cidade: %s", city)

	data, err := t.svc.GetWeatherByCityName(ctx, city)
	if err != nil {
		log.Printf("[MCP] Erro ao buscar clima: %v", err)
		return nil, nil, err
	}

	// Serializa struct domínio para JSON (bytes)
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("[MCP] Erro ao serializar resposta: %v", err)
		return nil, nil, err
	}

	result := &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(jsonBytes)},
		},
	}

	var responseMap map[string]any
	if err := json.Unmarshal(jsonBytes, &responseMap); err != nil {
		log.Printf("[MCP] não foi possível converter JSON para map: %v", err)
		return result, nil, nil
	}

	log.Printf("[MCP] Enviando resposta: %+v", responseMap)
	return result, responseMap, nil
}
