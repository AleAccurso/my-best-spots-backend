package usecases

type UsecaseBase struct {
	CategoryUsecase ICategoryUsecase
	SpotUsecase ISpotUsecase
}
