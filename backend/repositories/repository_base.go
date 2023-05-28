package repositories

type RepositoryBase struct {
	CategoryRepository ICategoryRepository
	AuthRepository     IAuthRepository
}
