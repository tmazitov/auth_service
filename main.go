package main

import (
	"github.com/tmazitov/auth_service.git/internal/handlers"
	service "github.com/tmazitov/auth_service.git/pkg/service"
)

func main() {

	var (
		auth *service.Service
	)

	auth = service.NewService("auth-service", "5000", "auth")
	auth.SetupHandlers(handlers.ServiceEndpoints())
	auth.Start()
}
