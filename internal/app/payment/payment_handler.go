package payment

import (
	"github.com/Sirpyerre/payment-platform/internal/service"
	"github.com/Sirpyerre/payment-platform/pkg/logger"
	"net/http"
	"strconv"

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

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Payment processed successfully",
	})
}

func (p *PaymentHandler) GetPayment(c echo.Context) error {
	transactionID, err := paramToInt(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid transaction ID")
	}

	// Lógica para obtener los detalles de un pago
	transaction, err := p.TransactionService.GetTransaction(transactionID)
	if err != nil {
		logger.GetLogger().Error("paymentHandler", "getPayment", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Error getting payment details",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, transaction)
}

func (p *PaymentHandler) ProcessRefund(c echo.Context) error {
	transactionID, err := paramToInt(c, "id")
	if err != nil && transactionID < 1 {
		return c.JSON(http.StatusBadRequest, "Invalid transaction ID")
	}

	err = p.TransactionService.RefundTransaction(transactionID)
	if err != nil {
		logger.GetLogger().Error("paymentHandler", "processRefund", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Error processing refund",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Refund successful",
		"status":  "Refunded",
	})
}

func paramToInt(c echo.Context, param string) (int, error) {
	id := c.Param(param)

	// convert transactionID to int
	transactionID, err := strconv.Atoi(id)
	if err != nil {
		logger.GetLogger().Error("paymentHandler", "paramToInt", err)
	}
	return transactionID, err
}
