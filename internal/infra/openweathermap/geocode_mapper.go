package openweathermap

import (
	"github.com/alvarolspeixoto/go-weather-mcp/internal/domain/geocode"
)

func GeocodeResponseToDomain(dto *OWGeocode) *geocode.Location {
	if dto == nil {
		return nil
	}

	return &geocode.Location{
		Latitude:  dto.Latitude,
		Longitude: dto.Longitude,
		City:      dto.Name,
		Country:   dto.Country,
	}
}
