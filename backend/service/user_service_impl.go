package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"unicode"

	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/repository"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
	Error          error
}

type Claims struct {
	Email  string
	UserID int
	Role   bool
	jwt.RegisteredClaims
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		fmt.Println("Validation error:", err)
		return web.UserResponse{} 
	}

	if !IsValidPassword(request.Password) {
		fmt.Println("Password does not meet security criteria")
		return web.UserResponse{Error: "Password must be at least 8 characters long and include an uppercase letter, a lowercase letter, a number, and a special character"}
	}

	tx, err := service.DB.Begin()
	if err != nil {
		fmt.Println("DB Begin error:", err)
		return web.UserResponse{} 
	}
	defer helper.CommitOrRollback(tx)

	hashedPassword, err := HashPassword(request.Password)
	if err != nil {
		fmt.Println("HashPassword error:", err)
		return web.UserResponse{} 
	}

	user := domain.User{
		Username:  request.Username,
		Password:  hashedPassword,
		Email:     request.Email,
		Role:      request.Role,
		NoTelepon: request.NoTelepon,
	}

	user = service.UserRepository.Register(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) (web.UserResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		return web.UserResponse{}, err
	}

	err = ComparePassword(user.Password, request.Password)
	if err != nil {
		return web.UserResponse{}, err
	}

	token, err := GenerateToken(user.Email, user.Role, user.Id, "yourSecretKey")
	if err != nil {
		return web.UserResponse{}, err
	}

	userResponse := helper.ToUserResponse(user)
	userResponse.Token = token

	return userResponse, nil
}

//utilities function >>>

func GenerateToken(email string, role bool, userId int, secretKey string) (string, error) {
	claims := &Claims{
		Email:  email,
		UserID: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secretKey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func IsValidPassword(password string) bool {
    var (
        hasMinLen  = false
        hasUpper   = false
        hasLower   = false
        hasNumber  = false
        hasSpecial = false
    )

    // Panjang minimal 8 karakter
    if len(password) >= 8 {
        hasMinLen = true
    }

    for _, char := range password {
        switch {
        case unicode.IsUpper(char):
            hasUpper = true
        case unicode.IsLower(char):
            hasLower = true
        case unicode.IsDigit(char):
            hasNumber = true
        case unicode.IsPunct(char) || unicode.IsSymbol(char):
            hasSpecial = true
        }
    }

    // Setidaknya harus ada satu karakter dari setiap kriteria
    return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
