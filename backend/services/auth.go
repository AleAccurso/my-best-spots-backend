package services

import (
	"my-best-spots-backend/repositories"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	repository repositories.Repository
}

type JWTClaim struct {
	Username      string `json:"Username"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	EmailVerified bool   `json:"email_verified"`
	jwt.StandardClaims
}

func InitialiseAuthService(repository repositories.Repository) AuthService {
	return AuthService{
		repository: repository,
	}
}

// func (service AuthService) Register(context *gin.Context, UserDTO dtos.UserReqCreateDTO) (*primitive.ObjectID, error) {

// 	User := mappers.UserReqCreateDTOToModel(UserDTO)

// 	if !utils.IsEmailValid(UserDTO.Email) {
// 		return nil, errors.New(constants.BAD_DATA + "email")
// 	}

// 	if !utils.IsAllowedLanguage(UserDTO.Language) {
// 		return nil, errors.New(constants.BAD_DATA + "language")
// 	}

// 	hashedPassword, err := service.getHash([]byte(UserDTO.Password))
// 	if err != nil {
// 		return nil, err
// 	}
// 	User.Password = *hashedPassword

// 	checkUser, _ := service.repository.UserRepository.GetUserByEmail(context, UserDTO.Email)
// 	if checkUser != nil {
// 		return nil, errors.New(constants.AUTH_EMAIL_EXISTS)
// 	}

// 	newId, err := service.repository.AuthRepository.AddUser(context, User)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return newId, nil
// }

// func (service AuthService) Login(context *gin.Context, loginReqDTO dtos.LoginReqDTO) (*string, error) {

// 	// check if email exists and password is correct
// 	User, err := service.repository.UserRepository.GetUserByEmail(context, loginReqDTO.Email)
// 	if err != nil {
// 		// c.IndentedJSON(http.StatusInternalServerError, err.Error())
// 		return nil, err
// 	}

// 	credentialError := service.checkPassword(User.Password, loginReqDTO.Password)
// 	if credentialError != nil {
// 		// context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
// 		// context.Abort()
// 		// return

// 		return nil, errors.New(constants.AUTH_PASSWORD_MISSMATCH)
// 	}

// 	tokenString, err := service.generateJWT(User.Nickname, User.Email, User.IsAdmin)
// 	if err != nil {
// 		// context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		// context.Abort()
// 		return nil, err
// 	}

// 	return &tokenString, nil
// }

// func (service AuthService) Logout(context *gin.Context) {
// 	cookie := http.Cookie{
// 		Name:   "token",
// 		MaxAge: -1}
// 	http.SetCookie(context.Writer, &cookie)
// }

// func (service AuthService) getHash(pwd []byte) (*string, error) {

// 	hash, err := bcrypt.GenerateFromPassword(pwd, 14)
// 	if err != nil {
// 		return nil, errors.New(constants.AUTH_UNABLE_TO_HASH_PASSWORD)
// 	}
// 	return lo.ToPtr(string(hash)), nil
// }

// func (service AuthService) checkPassword(UserPassword string, providedPassword string) error {
// 	err := bcrypt.CompareHashAndPassword([]byte(UserPassword), []byte(providedPassword))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (service AuthService) generateJWT(nickname string, email string, isAdmin bool) (string, error) {

// 	secretKey := os.Getenv("JWT_SECRET")

// 	expirationTime := time.Now().Add(24 * time.Hour)

// 	role := "User"

// 	if isAdmin {
// 		role = "admin"
// 	}

// 	claims := &JWTClaim{
// 		Email:         email,
// 		Username:  nickname,
// 		Role:          role,
// 		EmailVerified: true,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	tokenString, err := token.SignedString([]byte(secretKey))
// 	if err != nil {
// 		log.Println("Error in JWT token generation")
// 		return "", err
// 	}

// 	return tokenString, nil
// }
