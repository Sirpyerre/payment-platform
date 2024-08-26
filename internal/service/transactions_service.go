package service

import (
	"errors"
	"fmt"
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
	RefundRepository      repository.Refunder
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

func (s *TransactionService) RefundTransaction(transactionID int) error {
	// Get transaction
	transaction, err := s.TransactionRepository.GetTransaction(transactionID)
	if err != nil {
		return err
	}

	// Check if transaction is elegible for refund
	if transaction.Status != "success" {
		return errors.New("transaction is not elegible for refund")
	}

	// Process refund in bank
	responseRefund, err := s.BankService.RefundTransaction(transaction.TransactionBankID)
	if err != nil {
		return errors.New(fmt.Sprintf("error processing refund: %s", err.Error()))
	}

	// Check if refund was successful
	if responseRefund.Status != "success" {
		return errors.New("error processing refund")
	}

	// Save refund
	refund := new(models.Refund)
	refund.TransactionID = transactionID
	refund.Amount = transaction.Amount
	refund.Status = responseRefund.Status

	err = s.RefundRepository.Refund(refund)
	if err != nil {
		return errors.New(fmt.Sprintf("error saving refund: %s", err.Error()))
	}

	// Update transaction status
	err = s.TransactionRepository.UpdateTransactionStatus(transaction, "refunded")
	if err != nil {
		return errors.New(fmt.Sprintf("error updating transaction status: %s", err.Error()))
	}

	return nil
}
