package service

import "github.com/Sirpyerre/payment-platform/internal/models"

type UserServicer interface {
	Login(*models.LoginRequest) (string, error)
}
