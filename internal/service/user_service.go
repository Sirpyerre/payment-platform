package service

import (
	"errors"
	"time"

	"github.com/Sirpyerre/payment-platform/config"
	"github.com/Sirpyerre/payment-platform/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	Config *config.Configuration
}

func NewUserService(cfg *config.Configuration) *UserService {
	return &UserService{
		Config: cfg,
	}
}

func (s *UserService) Login(loginRequest *models.LoginRequest) (string, error) {
	if loginRequest.Username != "peter" || loginRequest.Password != "shhh!" {
		return "", errors.New("invalid credentials")
	}

	claims := &models.CustomClaims{
		Name:  "Jon Snow",
		Admin: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.Config.SecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
