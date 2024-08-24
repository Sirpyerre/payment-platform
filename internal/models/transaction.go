package models

import "time"

type TransactionsModel struct {
	ID                int       `json:"id"`
	MerchantID        int       `json:"merchant_id" validate:"required"`
	CustomerID        int       `json:"customer_id" validate:"required"`
	TransactionBankID string    `json:"transaction_bank_id"`
	Amount            float64   `json:"amount" validate:"required"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
