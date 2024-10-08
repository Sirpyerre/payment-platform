package repository

import (
	"context"
	"time"

	"github.com/Sirpyerre/payment-platform/config"
	"github.com/Sirpyerre/payment-platform/internal/dbconnection"
	"github.com/Sirpyerre/payment-platform/internal/models"
)

type TransactionRepository struct {
	Configuration *config.Configuration
	Connector     *dbconnection.Connector
	QueryTimeout  time.Duration
}

func NewTransactionRepository(cfg *config.Configuration, connector *dbconnection.Connector) *TransactionRepository {
	queryTimeout := time.Duration(cfg.DBConfig.QueryTimeout) * time.Second
	return &TransactionRepository{
		Configuration: cfg,
		Connector:     connector,
		QueryTimeout:  queryTimeout,
	}
}

func (t *TransactionRepository) Process(transaction *models.TransactionsModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), t.QueryTimeout)
	defer cancel()

	query := `INSERT INTO transactions (merchant_id, customer_id, amount, status,transaction_bank_id)
		VALUES ($1, $2, $3, $4, $5)
	`

	prepContext, err := t.Connector.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer prepContext.Close()

	_, err = prepContext.ExecContext(ctx, transaction.MerchantID, transaction.CustomerID,
		transaction.Amount, transaction.Status, transaction.TransactionBankID)
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepository) GetTransaction(transactionID int) (*models.TransactionsModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), t.QueryTimeout)
	defer cancel()

	query := `SELECT merchant_id, customer_id, amount, status, 
       transaction_bank_id, created_at
		FROM transactions
		WHERE id = $1
	`

	prepContext, err := t.Connector.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer prepContext.Close()

	transaction := new(models.TransactionsModel)
	transaction.ID = transactionID
	err = prepContext.QueryRowContext(ctx, transactionID).Scan(
		&transaction.MerchantID,
		&transaction.CustomerID,
		&transaction.Amount,
		&transaction.Status,
		&transaction.TransactionBankID,
		&transaction.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionRepository) UpdateTransactionStatus(transaction *models.TransactionsModel, status string) error {
	ctx, cancel := context.WithTimeout(context.Background(), t.QueryTimeout)
	defer cancel()

	updatedAt := time.Now()

	query := `UPDATE transactions SET status = $1, updated_at= $2 WHERE id = $3`
	prepContext, err := t.Connector.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer prepContext.Close()

	_, err = prepContext.ExecContext(ctx, status, updatedAt, transaction.ID)
	if err != nil {
		return err
	}

	return nil
}
