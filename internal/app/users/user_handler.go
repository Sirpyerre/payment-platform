package users

import (
	"github.com/Sirpyerre/payment-platform/internal/customvalidator"
	"github.com/Sirpyerre/payment-platform/internal/models"
	"github.com/Sirpyerre/payment-platform/internal/service"
	"github.com/Sirpyerre/payment-platform/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UsersHandler struct {
	UserService service.UserServicer
}

func NewUsersHandler(userService *service.UserService) *UsersHandler {
	return &UsersHandler{
		UserService: userService,
	}
}

func (u *UsersHandler) Login(c echo.Context) error {
	loginRequest := new(models.LoginRequest)

	if err := c.Bind(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error while binding request",
		})
	}

	c.Echo().Validator = &customvalidator.CustomValidator{Validator: validator.New()}
	if err := c.Validate(loginRequest); err != nil {
		logger.GetLogger().Error("Login", "error while validating login request", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error while validating login request",
			"error":   err.Error(),
		})
	}

	token, err := u.UserService.Login(loginRequest)
	if err != nil {
		logger.GetLogger().Error("Login", "error while logging in", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error while logging in",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
