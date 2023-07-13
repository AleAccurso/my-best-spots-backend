package repositories

type RepositoryBase struct {
	AddressRepository IAddressRepository
	AuthRepository     IAuthRepository
	CategoryRepository ICategoryRepository
	SpotRepository ISpotRepository
}
