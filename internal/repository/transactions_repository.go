package repository

import (
	"context"
	"github.com/Sirpyerre/payment-platform/config"
	"github.com/Sirpyerre/payment-platform/internal/dbconnection"
	"github.com/Sirpyerre/payment-platform/internal/models"
	"time"
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
       NULLIF(transaction_bank_id,0), created_at
		FROM transactions
		WHERE id = $1
	`

	prepContext, err := t.Connector.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer prepContext.Close()

	transaction := new(models.TransactionsModel)
	err = prepContext.QueryRowContext(ctx, transactionID).Scan(
		&transaction.ID,
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

	query := `UPDATE transactions SET status = $1 WHERE id = $2`
	prepContext, err := t.Connector.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer prepContext.Close()

	_, err = prepContext.ExecContext(ctx, transaction.ID, status)
	if err != nil {
		return err
	}

	return nil
}
