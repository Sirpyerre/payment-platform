package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Sirpyerre/payment-platform/config"
	"github.com/Sirpyerre/payment-platform/internal/models"
	"github.com/Sirpyerre/payment-platform/internal/repository"
	"github.com/Sirpyerre/payment-platform/pkg/httpcall"
	"github.com/Sirpyerre/payment-platform/pkg/logger"
	"time"
)

type TransactionService struct {
	TransactionRepository repository.Transactioner
	Config                *config.Configuration
	BankTimeout           time.Duration
}

func NewTransactionService(transactionRepository *repository.TransactionRepository, cfg *config.Configuration) *TransactionService {
	return &TransactionService{
		TransactionRepository: transactionRepository,
		Config:                cfg,
		BankTimeout:           time.Duration(cfg.BankProvider.Timeout) * time.Second,
	}
}

func (s *TransactionService) ProcessTransaction(transaction *models.TransactionsModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.BankTimeout)
	defer cancel()

	url := fmt.Sprintf("%s/payments", s.Config.BankProvider.URL)
	fmt.Println("url", url)
	response, err := httpcall.MakeCall(ctx, "POST", url, nil)
	if err != nil {
		logger.GetLogger().Error("TransactionService", "MakeCall", err)
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("error in banktransaction provider: %d", response.StatusCode)
	}

	bodyResponse := new(models.ResponsePayment)
	if err := json.NewDecoder(response.Body).Decode(bodyResponse); err != nil {
		return err
	}

	transaction.Status = bodyResponse.Status
	transaction.TransactionBankID = bodyResponse.TransactionID
	if err := s.TransactionRepository.Process(transaction); err != nil {
		return err
	}

	return nil

}
