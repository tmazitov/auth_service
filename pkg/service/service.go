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
		path    string
		handler Handler
		process []gin.HandlerFunc
	)

	for _, e := range endpoints {
		handler = e.Handler
		process = []gin.HandlerFunc{}
		process = append(process, handler.CoreBeforeMiddleware()...)
		process = append(process, handler.BeforeMiddleware()...)
		process = append(process, handler.Handle)
		process = append(process, handler.AfterMiddleware()...)
		process = append(process, handler.CoreAfterMiddleware()...)
		path = fmt.Sprintf("/%s/v0/api/%s", s.prefix, e.Path)
		s.core.Handle(e.Method, path, process...)
	}
}

func (s *Service) SetupDocs(endpoints []Endpoint) {

	var (
		path    string
		handler Handler
		process []gin.HandlerFunc
	)

	for _, e := range endpoints {
		handler = e.Handler
		process = []gin.HandlerFunc{}
		process = append(process, handler.CoreBeforeMiddleware()...)
		process = append(process, handler.BeforeMiddleware()...)
		process = append(process, handler.Handle)
		process = append(process, handler.AfterMiddleware()...)
		process = append(process, handler.CoreAfterMiddleware()...)
		path = e.Path
		s.core.Handle(e.Method, path, process...)
	}
}

func (s *Service) Start() {
	s.core.Run(":" + s.port)
}
