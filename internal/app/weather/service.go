package weather

import (
	"context"

	"github.com/alvarolspeixoto/go-weather-mcp/internal/domain/geocode"
	"github.com/alvarolspeixoto/go-weather-mcp/internal/domain/weather"
)

type Service struct {
	weatherUseCase *weather.UseCase
	geoUseCase     *geocode.UseCase
}

func NewService(weatherUC *weather.UseCase, geoUC *geocode.UseCase) *Service {
	return &Service{
		weatherUseCase: weatherUC,
		geoUseCase:     geoUC,
	}
}

func (s *Service) GetWeatherByCityName(ctx context.Context, cityName string) (*weather.Weather, error) {
	location, err := s.geoUseCase.GetCoordinatesByCityName(ctx, cityName)
	if err != nil {
		return nil, err
	}

	return s.weatherUseCase.GetWeatherByCoordinates(ctx, location.Latitude, location.Longitude)
}
