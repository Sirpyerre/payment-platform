package transactions

import (
	"net/http"

	"github.com/Sirpyerre/payment-platform/config"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	Config *config.Configuration
}

func NewTransactionHandler(cfg *config.Configuration) *TransactionHandler {
	return &TransactionHandler{
		Config: cfg,
	}
}

func (t *TransactionHandler) GetTransaction(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get transactions")
}

func (t *TransactionHandler) CreateTransaction(c echo.Context) error {
	return c.JSON(http.StatusOK, "Create transactions")
}

func (t *TransactionHandler) UpdateTransactionStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, "Update transactions status")
}
