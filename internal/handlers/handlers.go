package handlers

import "github.com/tmazitov/auth_service.git/pkg/service"

func ServiceEndpoints() []service.Endpoint {
	return []service.Endpoint{
		{Method: "GET", Path: "ping", Handler: PingHandler{}},
	}
}
