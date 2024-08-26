package httpserver

import (
	"github.com/Sirpyerre/payment-platform/internal/models"
	"github.com/Sirpyerre/payment-platform/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

func protectRoutesMiddleware(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		TokenLookup: "query:token,header:Authorization:Bearer ",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.CustomClaims)
		},
		SigningKey: []byte(secretKey),
		ErrorHandler: func(c echo.Context, err error) error {
			logger.GetLogger().Error("populateRoutes", "error while validating token", err)
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Unauthorized",
			})
		},
	})
}
