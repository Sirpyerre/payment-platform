package service

import "github.com/Sirpyerre/payment-platform/internal/models"

type TransactionsServicer interface {
	ProcessTransaction(*models.TransactionsModel) error
	GetTransaction(int) (*models.TransactionsModel, error)
}
