package openweathermap

import (
	"context"

	"github.com/alvarolspeixoto/go-weather-mcp/internal/domain/weather"
)

type WeatherRepository struct {
	client *Client
}

func NewWeatherRepository(c *Client) *WeatherRepository {
	return &WeatherRepository{client: c}
}

// Implementa o contrato do domínio
func (r *WeatherRepository) GetWeatherByCoordinates(ctx context.Context, lat, lon float64) (*weather.Weather, error) {
	dto, err := r.client.GetWeather(lat, lon)
	if err != nil {
		return nil, err
	}

	// Converte DTO → entity
	return WeatherResponseToDomain(dto), nil
}
