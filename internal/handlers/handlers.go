package handlers

import (
	"github.com/tmazitov/auth_service.git/internal/staff"
	"github.com/tmazitov/auth_service.git/pkg/service"
)

func ServiceEndpoints(st *staff.Staff) []service.Endpoint {
	return []service.Endpoint{
		{Method: "POST", Path: "ping", Handler: &PingHandler{}},
		{Method: "POST", Path: "verification/code", Handler: &CodeSendHandler{st: st}},
		{Method: "POST", Path: "verification/code/check", Handler: &CodeCheckHandler{st: st}},
		// {Method: "POST", Path: "token/refresh", Handler: &TokenRefreshHandler{st: st}},
	}
}
