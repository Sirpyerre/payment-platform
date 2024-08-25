package models

import (
	"time"
)

type TransactionsModel struct {
	ID                int       `json:"-"`
	MerchantID        int       `json:"merchant_id" validate:"required"`
	CustomerID        int       `json:"customer_id" validate:"required"`
	TransactionBankID int       `json:"transaction_bank_id"`
	Amount            float64   `json:"amount" validate:"required"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"-"`
}
