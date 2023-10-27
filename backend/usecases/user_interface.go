package usecases

import (
	"my-best-spots-backend/dtos"

	"github.com/gin-gonic/gin"
)

type IUserUsecase interface {
	GetUserByEmail(c *gin.Context, email string) (*dtos.UserResDTO, error)
}
