package repository

import "github.com/Sirpyerre/payment-platform/internal/models"

type Refunder interface {
	Refund(*models.Refund) error
}
