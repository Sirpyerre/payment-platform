package payment

import (
	"github.com/Sirpyerre/payment-platform/internal/service"
	"github.com/Sirpyerre/payment-platform/pkg/logger"
	"net/http"

	"github.com/Sirpyerre/payment-platform/internal/customvalidator"
	"github.com/Sirpyerre/payment-platform/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	Validator          echo.Validator
	TransactionService service.TransactionsServicer
}

func NewPaymentHandler(transactionService *service.TransactionService) *PaymentHandler {
	return &PaymentHandler{
		Validator:          &customvalidator.CustomValidator{Validator: validator.New()},
		TransactionService: transactionService,
	}
}

func (p *PaymentHandler) ProcessPayment(c echo.Context) error {
	transaction := new(models.TransactionsModel)
	if err := c.Bind(transaction); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	c.Echo().Validator = &customvalidator.CustomValidator{Validator: validator.New()}
	if err := c.Validate(transaction); err != nil {
		logger.GetLogger().Error("paymentHandler", "proccessPayment", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// Lógica para procesar un pago
	err := p.TransactionService.ProcessTransaction(transaction)
	if err != nil {
		logger.GetLogger().Error("paymentHandler", "proccessPayment", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Error processing payment",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, "Payment processed")
}

func (p *PaymentHandler) GetPayment(c echo.Context) error {
	// Lógica para obtener detalles de un pago
	return c.JSON(http.StatusOK, "Payment details")
}

func (p *PaymentHandler) ProcessRefund(c echo.Context) error {
	// Lógica para procesar un reembolso
	return c.JSON(http.StatusOK, "Refund processed")
}
