package service

import "github.com/Sirpyerre/payment-platform/internal/models"

type TransactionsServicer interface {
	ProcessTransaction(model *models.TransactionsModel) error
}
