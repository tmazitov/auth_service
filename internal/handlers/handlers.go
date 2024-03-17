package handlers

import (
	"github.com/tmazitov/auth_service.git/internal/staff"
	"github.com/tmazitov/auth_service.git/pkg/service"
)

func ServiceEndpoints(st *staff.Staff) []service.Endpoint {
	return []service.Endpoint{
		{Method: "GET", Path: "ping", Handler: &PingHandler{}},
		{Method: "POST", Path: "verification/code", Handler: &CodeSendHandler{st: st}},
		{Method: "POST", Path: "verification/code/check", Handler: &CodeCheckHandler{st: st}},
		{Method: "POST", Path: "token/refresh", Handler: &TokenRefreshHandler{st: st}},
		{Method: "GET", Path: "/oauth/google/login", Handler: &OauthGoogleLoginHandler{st: st}},
		{Method: "GET", Path: "/oauth/google/callback", Handler: &OauthGoogleCallbackHandler{st: st}},
	}
}

func ServiceDocs() []service.Endpoint {
	return []service.Endpoint{
		{Method: "GET", Path: "/swagger/*any", Handler: &SwaggerDocsHandler{}},
	}
}
