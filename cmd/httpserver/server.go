package httpserver

import (
	"fmt"
	"go.uber.org/dig"
	"net/http"
	"time"

	"github.com/Sirpyerre/payment-platform/config"
	"github.com/Sirpyerre/payment-platform/internal/di"
	"github.com/Sirpyerre/payment-platform/pkg/logger"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type Server struct {
	server       *echo.Echo
	allowOrigins []string
	allowMethods []string
	allowHeaders []string
	logger       *logger.Log
}

// NewServer :
func NewServer(log *logger.Log) *Server {
	return &Server{
		server:       echo.New(),
		allowOrigins: []string{"*"},
		allowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		allowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
		logger:       log,
	}
}

func (s *Server) Init() {
	s.server.HideBanner = true
	container := di.GetContainer()
	cfg := s.getConfiguration(container)

	s.setupMiddlewares()
	s.setupHandlers()
	s.setupRoutes(container, cfg)
	s.startServer(cfg)
}

func (s *Server) getConfiguration(container *dig.Container) *config.Configuration {
	cfg := new(config.Configuration)
	err := container.Invoke(func(configuration *config.Configuration) {
		cfg = configuration
	})
	s.logger.FatalIfError("server", "getConfiguration", err)
	return cfg
}

func (s *Server) setupMiddlewares() {
	s.server.Use(
		echoMiddleware.Recover(),
		echoMiddleware.RequestID(),
		echoMiddleware.GzipWithConfig(echoMiddleware.GzipConfig{
			Level: 5,
		}),
		echoMiddleware.TimeoutWithConfig(echoMiddleware.TimeoutConfig{
			Timeout: 40 * time.Second,
		}),
		s.buildCORSMiddleware(s.allowOrigins, s.allowMethods, s.allowHeaders),
		RequestLogger,
	)
}

func (s *Server) setupHandlers() {
	s.setupNotFoundHandler()
	s.setupErrorHandler()
}

func (s *Server) setupRoutes(container *dig.Container, cfg *config.Configuration) {
	RegisterRoutes(container, cfg)
	for _, r := range ServerRoutes {
		s.server.Add(r.Method, r.Pattern, r.HandlerFunc, r.MiddleWares...).Name = r.Name
	}
}

func (s *Server) startServer(cfg *config.Configuration) {
	s.logger.Fatal("server", "startServer", s.server.Start(fmt.Sprintf(":%d", cfg.Server.Port)))
}

func (s *Server) buildCORSMiddleware(allowOrigins, allowMethods, allowHeaders []string) echo.MiddlewareFunc {
	corsConfig := echoMiddleware.CORSConfig{
		AllowOrigins: allowOrigins,
		AllowMethods: allowMethods,
		AllowHeaders: allowHeaders,
	}
	return echoMiddleware.CORSWithConfig(corsConfig)
}

// setupNotFoundHandler :
func (s *Server) setupNotFoundHandler() {
	echo.NotFoundHandler = func(c echo.Context) error {
		httpError := echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("resource not found :%s", c.Request().URL))
		return c.JSON(httpError.Code, httpError)
	}
}

// setupErrorHandler :
func (s *Server) setupErrorHandler() {
	s.server.HTTPErrorHandler = func(err error, context echo.Context) {
		s.logger.Request(context, time.Now())
		s.logger.Error("server", "setupErrorHandler", err)

		_ = context.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
}
