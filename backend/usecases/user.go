package usecases

import (
	"my-best-spots-backend/dtos"
	"my-best-spots-backend/repositories"

	"github.com/gin-gonic/gin"
)

type UserUsecase struct {
	repositories repositories.Repository
}

func InitialiseUserUsecase(repositories repositories.Repository) UserUsecase {
	return UserUsecase{
		repositories: repositories,
	}
}

func (usecase UserUsecase) GetUserByEmail(c *gin.Context, email string) (*dtos.UserResDTO, error) {
	return nil, nil
}
