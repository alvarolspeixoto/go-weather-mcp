package openweathermap

import (
	"context"

	"github.com/alvarolspeixoto/go-weather-mcp/internal/domain/geocode"
)

type GeocodeRepository struct {
	client *Client
}

func NewGeocodeRepository(c *Client) *GeocodeRepository {
	return &GeocodeRepository{client: c}
}

func (r *GeocodeRepository) GetCoordinatesByCityName(ctx context.Context, city string) (*geocode.Location, error) {
	dto, err := r.client.GetGeocode(city)
	if err != nil {
		return nil, err
	}

	return GeocodeResponseToDomain(dto), nil
}
