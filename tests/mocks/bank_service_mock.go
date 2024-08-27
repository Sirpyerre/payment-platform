package mocks

import (
	"github.com/Sirpyerre/payment-platform/internal/models"
	"github.com/stretchr/testify/mock"
)

type BankServiceMock struct {
	mock.Mock
}

func (m *BankServiceMock) ProcessTransaction() (*models.ResponsePayment, error) {
	args := m.Called()
	return args.Get(0).(*models.ResponsePayment), args.Error(1)
}

func (m *BankServiceMock) RefundTransaction(transactionID int) (*models.ResponsePayment, error) {
	args := m.Called(transactionID)
	return args.Get(0).(*models.ResponsePayment), args.Error(1)
}
