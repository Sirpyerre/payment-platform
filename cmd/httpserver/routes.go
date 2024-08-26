package httpserver

import (
	"github.com/Sirpyerre/payment-platform/config"
	"github.com/Sirpyerre/payment-platform/internal/app/payment"
	"github.com/Sirpyerre/payment-platform/internal/app/transactions"
	"github.com/Sirpyerre/payment-platform/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type Routes []Route

// Route :
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc echo.HandlerFunc
	MiddleWares []echo.MiddlewareFunc
}

var ServerRoutes Routes

func RegisterRoutes(container *dig.Container, cfg *config.Configuration) {
	errInvoke := container.Invoke(func(
		paymentHandler *payment.PaymentHandler,
		transactionHandler *transactions.TransactionHandler,
	) {
		ServerRoutes = setRoutes(paymentHandler, transactionHandler, cfg)
	})

	logger.GetLogger().FatalIfError("httpserver", "RegisterRoutes", errInvoke)
}

func setRoutes(paymentHandler *payment.PaymentHandler, transaction *transactions.TransactionHandler, _ *config.Configuration) Routes {
	return Routes{
		Route{
			"ProcessPayment",
			"POST",
			"/api/v1/payment/process",
			paymentHandler.ProcessPayment,
			nil,
		},
		Route{
			"GetPayment",
			"GET",
			"/api/v1/payment/:id",
			paymentHandler.GetPayment,
			nil,
		},
		Route{
			"ProcessRefund",
			"POST",
			"/api/v1/payment/:id/refund",
			paymentHandler.ProcessRefund,
			nil,
		},
	}
}
