package weather

import "context"

type Repository interface {
	GetWeatherByCoordinates(ctx context.Context, lat, lon float64) (*Weather, error)
}
