package geocode

import "context"

type Repository interface {
	GetCoordinatesByCityName(ctx context.Context, cityName string) (*Location, error)
}
