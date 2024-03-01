package handlers

import (
	"github.com/tmazitov/auth_service.git/internal/staff"
	"github.com/tmazitov/auth_service.git/pkg/service"
)

func ServiceEndpoints(st *staff.Staff) []service.Endpoint {
	return []service.Endpoint{
		// {Method: "GET", Path: "ping", Handler: PingHandler{}},
		{Method: "POST", Path: "send-code", Handler: SendCodeHandler{st: st}},
	}
}
