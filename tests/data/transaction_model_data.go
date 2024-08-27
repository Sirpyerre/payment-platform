package data

import (
	"github.com/Sirpyerre/payment-platform/internal/models"
	"time"
)

var TransactionModelData = &models.TransactionsModel{
	ID:                1,
	MerchantID:        2,
	CustomerID:        2,
	TransactionBankID: 12345657890,
	Amount:            100.10,
	Status:            "",
	CreatedAt:         time.Now(),
	UpdatedAt:         time.Now(),
}
