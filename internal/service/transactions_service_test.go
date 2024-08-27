package service

import (
	"github.com/Sirpyerre/payment-platform/config"
	"github.com/Sirpyerre/payment-platform/tests/data"
	"github.com/Sirpyerre/payment-platform/tests/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTransactionService(t *testing.T) {
	cfg := new(config.Configuration)
	cfg.BankProvider.Timeout = 10

	t.Run("Test ProcessTransaction Success", func(t *testing.T) {
		mockTransactionRepo := new(mocks.TransactionRepositoryMock)
		mockBankService := new(mocks.BankServiceMock)
		mockRefundRepo := new(mocks.RefundRepositoryMock)
		transactionService := TransactionService{
			TransactionRepository: mockTransactionRepo,
			BankService:           mockBankService,
			RefundRepository:      mockRefundRepo,
			Config:                cfg,
			BankTimeout:           time.Duration(cfg.BankProvider.Timeout) * time.Second,
		}

		mockBankService.On("ProcessTransaction").Return(data.ResponsePaymentData, nil)
		mockTransactionRepo.On("Process", data.TransactionModelData).Return(nil)

		err := transactionService.ProcessTransaction(data.TransactionModelData)
		assert.Nil(t, err)
	})
}
