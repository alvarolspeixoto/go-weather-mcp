package geocode

import (
	"github.com/alvarolspeixoto/go-weather-mcp/internal/domain/geocode"
	"github.com/alvarolspeixoto/go-weather-mcp/internal/infra/openweathermap"
)

func ToDomain(dtos []openweathermap.OWGeocode) *geocode.Location {
	if len(dtos) == 0 {
		return nil
	}
	dto := dtos[0]
	return &geocode.Location{
		Latitude:  dto.Latitude,
		Longitude: dto.Longitude,
		City:      dto.Name,
		Country:   dto.Country,
	}
}
