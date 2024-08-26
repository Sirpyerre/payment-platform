package httpserver

import (
	"github.com/labstack/echo/v4"

	"github.com/Sirpyerre/payment-platform/config"
	"github.com/Sirpyerre/payment-platform/internal/app/payment"
	"github.com/Sirpyerre/payment-platform/internal/app/users"
	"github.com/Sirpyerre/payment-platform/pkg/logger"

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
		usersHandler *users.UsersHandler,
	) {
		ServerRoutes = setRoutes(paymentHandler, usersHandler, cfg)
	})

	logger.GetLogger().FatalIfError("httpserver", "RegisterRoutes", errInvoke)
}

func setRoutes(paymentHandler *payment.PaymentHandler, usersHandler *users.UsersHandler, cfg *config.Configuration) Routes {
	jwtConfig := protectRoutesMiddleware(cfg.SecretKey)
	middlewares := []echo.MiddlewareFunc{
		jwtConfig,
	}

	return Routes{
		Route{
			"Login",
			"POST",
			"/api/v1/login",
			usersHandler.Login,
			nil,
		},
		Route{
			"ProcessPayment",
			"POST",
			"/api/v1/payment/process",
			paymentHandler.ProcessPayment,
			middlewares,
		},
		Route{
			"GetPayment",
			"GET",
			"/api/v1/payment/:id",
			paymentHandler.GetPayment,
			middlewares,
		},
		Route{
			"ProcessRefund",
			"POST",
			"/api/v1/payment/:id/refund",
			paymentHandler.ProcessRefund,
			middlewares,
		},
	}
}
