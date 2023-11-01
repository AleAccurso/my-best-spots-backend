package repositories

type RepositoryBase struct {
	AddressRepository  IAddressRepository
	CategoryRepository ICategoryRepository
	SpotRepository     ISpotRepository
}
