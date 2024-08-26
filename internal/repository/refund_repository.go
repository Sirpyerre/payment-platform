package repository

import (
	"context"
	"github.com/Sirpyerre/payment-platform/config"
	"github.com/Sirpyerre/payment-platform/internal/dbconnection"
	"github.com/Sirpyerre/payment-platform/internal/models"
	"time"
)

type RefundRepository struct {
	Configuration *config.Configuration
	Connector     *dbconnection.Connector
	QueryTimeout  time.Duration
}

func NewRefundRepository(cfg *config.Configuration, connector *dbconnection.Connector) *RefundRepository {
	queryTimeout := time.Duration(cfg.DBConfig.QueryTimeout) * time.Second
	return &RefundRepository{
		Configuration: cfg,
		Connector:     connector,
		QueryTimeout:  queryTimeout,
	}
}

func (r *RefundRepository) Refund(refund *models.Refund) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.QueryTimeout)
	defer cancel()

	query := `INSERT INTO refunds (transaction_id, amount, status) VALUES ($1, $2, $3)`
	prepContext, err := r.Connector.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer prepContext.Close()

	_, err = prepContext.ExecContext(ctx, refund.TransactionID, refund.Amount, refund.Status)
	if err != nil {
		return err
	}

	return nil
}
