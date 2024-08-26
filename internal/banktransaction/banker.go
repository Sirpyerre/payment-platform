package banktransaction

import "github.com/Sirpyerre/payment-platform/internal/models"

type Banker interface {
	ProcessTransaction() (*models.ResponsePayment, error)
	RefundTransaction(int) (*models.ResponsePayment, error)
}
