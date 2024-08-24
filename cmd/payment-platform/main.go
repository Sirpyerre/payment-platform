package main

import (
	"github.com/Sirpyerre/payment-platform/cmd/httpserver"
	"github.com/Sirpyerre/payment-platform/pkg/logger"
)

func main() {
	server := httpserver.NewServer(logger.GetLogger())
	server.Init()
}
