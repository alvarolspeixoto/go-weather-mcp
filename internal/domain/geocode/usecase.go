package geocode

import "context"

type UseCase struct {
	repository Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{
		repository: repo,
	}
}

func (uc *UseCase) GetCoordinatesByCityName(ctx context.Context, cityName string) (*Location, error) {
	return uc.repository.GetCoordinatesByCityName(ctx, cityName)
}
