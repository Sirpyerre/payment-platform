package mocks

import (
	"github.com/Sirpyerre/payment-platform/internal/models"
	"github.com/stretchr/testify/mock"
)

type TransactionServiceMock struct {
	mock.Mock
}

func (m *TransactionServiceMock) ProcessTransaction(transaction *models.TransactionsModel) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *TransactionServiceMock) GetTransaction(transactionID int) (*models.TransactionsModel, error) {
	args := m.Called(transactionID)
	return args.Get(0).(*models.TransactionsModel), args.Error(1)
}

func (m *TransactionServiceMock) RefundTransaction(transactionID int) error {
	args := m.Called(transactionID)
	return args.Error(0)
}
