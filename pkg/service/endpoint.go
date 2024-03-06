package service

type Endpoint struct {
	Method  string
	Path    string
	Handler Handler
}
