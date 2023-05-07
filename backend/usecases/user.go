package usecases

import (
	"context"
	"errors"
	"math"
	"new-rating-movies-go-backend/constants"
	"new-rating-movies-go-backend/dtos"
	"new-rating-movies-go-backend/helpers"
	"new-rating-movies-go-backend/repositories"
	"new-rating-movies-go-backend/usecases/mappers"
	"new-rating-movies-go-backend/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repositories repositories.Repository
}

func InitialiseUserUsecase(repositories repositories.Repository) UserUsecase {
	return UserUsecase{
		repositories: repositories,
	}
}

func (usecase UserUsecase) GetUsers(c *gin.Context, page string, size string) (*dtos.UserPagingResDTO, error) {

	ctx := context.TODO()

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + page)
	}
	if pageInt <= 0 {
		pageInt = 1
	}

	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + size)
	}
	if sizeInt <= 0 {
		sizeInt = 8
	}

	usersCount, err := usecase.repositories.UserRepository.CountUsers(ctx)
	if err != nil {
		return nil, err
	}

	nbPages := math.Ceil(float64(*usersCount) / float64(sizeInt))

	if nbPages == 0 {
		nbPages = 1
	}

	if float64(pageInt) > nbPages-1 {
		pageInt = int(nbPages)
	}

	pagingUsers := dtos.UserPagingResDTO{
		Page:      int8(pageInt),
		Size:      int8(sizeInt),
		NbPages:   int8(nbPages),
		NbResults: int16(*usersCount),
	}

	users, err := usecase.repositories.UserRepository.GetUsers(ctx, pageInt, sizeInt)
	if err != nil {
		return nil, err
	}

	pagingUsers.Data = mappers.UserModelsToResDTOs(users)

	return &pagingUsers, nil
}

func (usecase UserUsecase) GetUserById(c *gin.Context, userId string) (*dtos.UserResDTO, error) {

	ctx := context.TODO()

	userUUID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "userId")
	}

	user, err := usecase.repositories.UserRepository.GetUserById(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	err = helpers.IsLoggedUserOrAdmin(c, userUUID, *user)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.UserModelToResDTO(*user)

	return &userDTO, nil
}

func (usecase UserUsecase) GetUserByEmail(c *gin.Context, email string) (*dtos.UserResDTO, error) {

	ctx := context.TODO()

	loggedUserEmail, ok := c.Get("user_email")
	if !ok {
		return nil, errors.New(constants.AUTH_UNVERIFIED_EMAIL)
	}

	loggedUserRole, ok := c.Get("user_role")
	if !ok {
		return nil, errors.New("cannot get logged user role")
	}

	if !utils.IsEmailValid(email) {
		return nil, errors.New(constants.BAD_DATA + "email")
	}

	user, err := usecase.repositories.UserRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user.Email != loggedUserEmail && loggedUserRole != "admin" {
		return nil, errors.New(constants.AUTH_UNAUTHORIZED)
	}

	userDTO := mappers.UserModelToResDTO(*user)

	return &userDTO, nil
}

func (usecase UserUsecase) UpdateUserById(c *gin.Context, userId string, reqUpdateDTO dtos.UserReqUpdateDTO) (*dtos.UserResDTO, error) {

	ctx := context.TODO()

	userUUID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "userId")
	}

	userNewInfo := mappers.UserReqUpdateDTOToModel(reqUpdateDTO)

	existinguser, err := usecase.repositories.UserRepository.GetUserById(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	err = helpers.IsLoggedUserOrAdmin(c, userUUID, *existinguser)
	if err != nil {
		return nil, err
	}

	userNewInfo.Id = existinguser.Id
	userNewInfo.CreatedAt = existinguser.CreatedAt
	userNewInfo.Password = existinguser.Password

	if reqUpdateDTO.Password != "" {
		hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(reqUpdateDTO.Password), 14)
		if err != nil {
			return nil, errors.New(constants.AUTH_UNABLE_TO_HASH_PASSWORD)
		}
		userNewInfo.Password = string(hashedNewPassword)
	}

	err = usecase.repositories.UserRepository.ModifyUserById(ctx, userNewInfo)
	if err != nil {
		return nil, err
	}

	updatedUser, err := usecase.repositories.UserRepository.GetUserById(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	if reqUpdateDTO.Nickname != updatedUser.Nickname || reqUpdateDTO.Email != updatedUser.Email || reqUpdateDTO.Admin != updatedUser.IsAdmin || reqUpdateDTO.Language != updatedUser.Language || reqUpdateDTO.ProfilePic != updatedUser.ProfilePic {
		return nil, errors.New("something whent wrong during the update")
	}

	userResDTO := mappers.UserModelToResDTO(*updatedUser)

	return &userResDTO, nil
}

func (usecase UserUsecase) DeleteUserById(c *gin.Context, userId string) (*primitive.ObjectID, error) {

	ctx := context.TODO()

	userUUID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New(constants.BAD_PARAMS + "userId")
	}

	user, err := usecase.repositories.UserRepository.GetUserById(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	err = helpers.IsLoggedUserOrAdmin(c, userUUID, *user)
	if err != nil {
		return nil, err
	}

	err = usecase.repositories.UserRepository.DeleteUserById(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	return &user.Id, nil
}