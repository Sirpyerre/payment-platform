package service

import (
	"github.com/Sirpyerre/payment-platform/config"
	"github.com/Sirpyerre/payment-platform/internal/banktransaction"
	"github.com/Sirpyerre/payment-platform/internal/models"
	"github.com/Sirpyerre/payment-platform/internal/repository"
	"strconv"
	"time"
)

type TransactionService struct {
	TransactionRepository repository.Transactioner
	BankService           banktransaction.Banker
	Config                *config.Configuration
	BankTimeout           time.Duration
}

func NewTransactionService(transactionRepository *repository.TransactionRepository, bankService *banktransaction.BankTransaction,
	cfg *config.Configuration) *TransactionService {
	return &TransactionService{
		TransactionRepository: transactionRepository,
		Config:                cfg,
		BankTimeout:           time.Duration(cfg.BankProvider.Timeout) * time.Second,
		BankService:           bankService,
	}
}

func (s *TransactionService) ProcessTransaction(transaction *models.TransactionsModel) error {
	responsePayment, err := s.BankService.ProcessTransaction()
	if err != nil {
		return err
	}

	transactionBankID, err := strconv.Atoi(responsePayment.TransactionID)
	if err != nil {
		return err
	}

	transaction.Status = responsePayment.Status
	transaction.TransactionBankID = transactionBankID
	if err := s.TransactionRepository.Process(transaction); err != nil {
		return err
	}

	return nil
}

func (s *TransactionService) GetTransaction(transactionID int) (*models.TransactionsModel, error) {
	transaction, err := s.TransactionRepository.GetTransaction(transactionID)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
