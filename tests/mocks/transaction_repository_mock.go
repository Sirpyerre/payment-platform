package mocks

import (
	"github.com/Sirpyerre/payment-platform/internal/models"
	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (m *TransactionRepositoryMock) Process(transaction *models.TransactionsModel) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) GetTransaction(transactionID int) (*models.TransactionsModel, error) {
	args := m.Called(transactionID)
	return args.Get(0).(*models.TransactionsModel), args.Error(1)
}

func (m *TransactionRepositoryMock) UpdateTransactionStatus(transaction *models.TransactionsModel, status string) error {
	args := m.Called(transaction, status)
	return args.Error(0)
}
