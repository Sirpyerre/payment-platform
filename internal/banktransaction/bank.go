package banktransaction

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Sirpyerre/payment-platform/config"
	"github.com/Sirpyerre/payment-platform/internal/models"
	"github.com/Sirpyerre/payment-platform/pkg/httpcall"
	"github.com/Sirpyerre/payment-platform/pkg/logger"
	"time"
)

type BankTransaction struct {
	Config  *config.Configuration
	Timeout time.Duration
}

func NewBankTransaction(cfg *config.Configuration) *BankTransaction {
	timeout := time.Duration(cfg.BankProvider.Timeout) * time.Second
	return &BankTransaction{
		Config:  cfg,
		Timeout: timeout,
	}
}

func (b *BankTransaction) ProcessTransaction() (*models.ResponsePayment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), b.Timeout)
	defer cancel()

	url := fmt.Sprintf("%s/payments", b.Config.BankProvider.URL)
	fmt.Println("url", url)
	response, err := httpcall.MakeCall(ctx, "POST", url, nil)
	if err != nil {
		logger.GetLogger().Error("TransactionService", "MakeCall", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error in banktransaction provider: %d", response.StatusCode)
	}

	responsePayment := new(models.ResponsePayment)
	if err := json.NewDecoder(response.Body).Decode(responsePayment); err != nil {
		return nil, err
	}

	return responsePayment, nil
}

func (b *BankTransaction) RefundTransaction(transactionBankID int) (*models.ResponsePayment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), b.Timeout)
	defer cancel()

	url := fmt.Sprintf("%s/payments/%d/refund", b.Config.BankProvider.URL, transactionBankID)
	response, err := httpcall.MakeCall(ctx, "POST", url, nil)
	if err != nil {
		logger.GetLogger().Error("TransactionService", "MakeCall", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error in banktransaction provider: %d", response.StatusCode)
	}

	responsePayment := new(models.ResponsePayment)
	if err := json.NewDecoder(response.Body).Decode(responsePayment); err != nil {
		return nil, err
	}

	return responsePayment, nil
}
