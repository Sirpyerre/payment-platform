package repository

import "github.com/Sirpyerre/payment-platform/internal/models"

type Transactioner interface {
	Process(model *models.TransactionsModel) error
}
