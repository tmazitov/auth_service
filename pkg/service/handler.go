package service

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Handle(ctx *gin.Context)
	AfterMiddleware() []gin.HandlerFunc
	BeforeMiddleware() []gin.HandlerFunc
	CoreAfterMiddleware() []gin.HandlerFunc
	CoreBeforeMiddleware() []gin.HandlerFunc
}

type HandlerInput any
type HandlerOutput any

type HandlerCoreBehavior[I HandlerInput, O HandlerOutput] struct {
	Input  I
	Output O
}

func (h *HandlerCoreBehavior[I, O]) ReadInput(ctx *gin.Context) error {
	if ctx.Request.Method == "GET" {
		return ctx.ShouldBindQuery(&h.Input)
	}
	return ctx.ShouldBindJSON(&h.Input)
}

/*
BeforeMiddleware return the list of your middleware functions that
will be executed before Handler func. Implement as you want
*/
func (h *HandlerCoreBehavior[I, O]) BeforeMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

/*
	DO NOT EDIT THIS METHOD

CoreBeforeMiddleware return a basic list of middleware functions that will be
executed before Handler and BeforeMiddleware function list
Contain :
  - Automatic output writer
*/
func (h *HandlerCoreBehavior[I, O]) CoreBeforeMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		h.readInputMiddleware(),
	}
}

/*
AfterMiddleware return the list of your middleware functions that
will be executed after Handler func. Implement as you want
*/
func (h *HandlerCoreBehavior[I, O]) AfterMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

/*
	DO NOT EDIT THIS METHOD

CoreAfterMiddleware return a basic list of middleware functions that will be
executed after Handler and AfterMiddleware function list
Contain :
  - Automatic output writer
*/
func (h *HandlerCoreBehavior[I, O]) CoreAfterMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		h.writeOutputMiddleware(),
	}
}

type HandlerMonoWriteBehavior[O HandlerOutput] struct {
	HandlerCoreBehavior[struct{}, O]
}

type HandlerMonoReadBehavior[I HandlerInput] struct {
	HandlerCoreBehavior[I, struct{}]
}

type HandlerClearBehavior struct {
	HandlerCoreBehavior[struct{}, struct{}]
}

func (h *HandlerClearBehavior) CoreBeforeMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

func (h *HandlerClearBehavior) CoreAfterMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}
