package services

import "my-best-spots-backend/repositories"

type Service struct {
	ServiceBase
}

func Initialise(repository repositories.Repository) Service {
	return Service{
		ServiceBase: ServiceBase{
			AuthService: InitialiseAuthService(repository),
		},
	}
}
