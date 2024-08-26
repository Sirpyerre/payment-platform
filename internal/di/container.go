package di

import (
	"github.com/Sirpyerre/payment-platform/config"
	"github.com/Sirpyerre/payment-platform/internal/app/payment"
	"github.com/Sirpyerre/payment-platform/internal/app/transactions"
	"github.com/Sirpyerre/payment-platform/internal/banktransaction"
	"github.com/Sirpyerre/payment-platform/internal/dbconnection"
	"github.com/Sirpyerre/payment-platform/internal/repository"
	"github.com/Sirpyerre/payment-platform/internal/service"
	"github.com/Sirpyerre/payment-platform/pkg/logger"
	"go.uber.org/dig"

	"sync"
)

var (
	container *dig.Container
	once      sync.Once
)

// GetContainer :
func GetContainer() *dig.Container {
	once.Do(func() {
		container = buildContainer()
	})
	return container
}

func buildContainer() *dig.Container {
	c := dig.New()

	logger.GetLogger().FatalIfError("di", "buildContainer",
		c.Provide(logger.NewLog),
		c.Provide(config.NewConfiguration),
		c.Provide(dbconnection.NewDBConnection),
		c.Provide(banktransaction.NewBankTransaction),
		c.Provide(repository.NewRefundRepository),
		c.Provide(repository.NewTransactionRepository),
		c.Provide(service.NewTransactionService),
		c.Provide(payment.NewPaymentHandler),
		c.Provide(transactions.NewTransactionHandler),
	)

	return c

}
