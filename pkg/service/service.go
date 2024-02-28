package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Service struct {
	port   string
	name   string
	prefix string
	core   *gin.Engine
}

func NewService(name string, port string, prefix string) *Service {
	return &Service{
		port:   port,
		name:   name,
		prefix: prefix,
		core:   gin.Default(),
	}
}

func (s *Service) SetupMiddleware(middlewares []gin.HandlerFunc) {
	for _, middleware := range middlewares {
		s.core.Use(middleware)
	}
}

func (s *Service) SetupHandlers(endpoints []Endpoint) {

	var (
		handler Handler
		process []gin.HandlerFunc
	)

	for _, e := range endpoints {
		handler = e.Handler
		process = []gin.HandlerFunc{}
		process = append(process, handler.Middleware()...)
		process = append(process, handler.Handle)
		s.core.Handle(e.Method, fmt.Sprintf("/%s/v0/api/%s", s.prefix, e.Path), process...)
	}
}

func (s *Service) Start() {
	s.core.Run(":" + s.port)
}
