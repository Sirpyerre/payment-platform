package httpserver

import (
	"time"

	"github.com/Sirpyerre/payment-platform/pkg/logger"

	"github.com/labstack/echo/v4"
)

// RequestLogger :
func RequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) (e error) {
		start := time.Now()
		if e := next(context); e != nil {
			return e
		}
		logger.GetLogger().Request(context, start)
		return
	}
}
