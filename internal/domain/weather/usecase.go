package weather

import "context"

type UseCase struct {
	repository Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{
		repository: repo,
	}
}

func (uc *UseCase) GetWeatherByCoordinates(ctx context.Context, lat, lon float64) (*Weather, error) {
	return uc.repository.GetWeatherByCoordinates(ctx, lat, lon)
}
