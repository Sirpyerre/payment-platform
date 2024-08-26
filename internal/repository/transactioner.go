package repository

import "github.com/Sirpyerre/payment-platform/internal/models"

type Transactioner interface {
	Process(*models.TransactionsModel) error
	GetTransaction(int) (*models.TransactionsModel, error)
	UpdateTransactionStatus(transaction *models.TransactionsModel, status string) error
}
