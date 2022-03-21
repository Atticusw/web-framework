package web

import (
	"net/http"
)

type HandlerFunc func(*Context)

// map k--> url and v --> handler
type Engine struct {
	router *router
}

// return router map
func New() *Engine {
	return &Engine{router: newRouter()}
}

// add route
func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

// GET defines the method to add request
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

// POST defines the method to add request
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

// implement ServeHTTP and custom handle method
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.handle(c)
}

// Run http server
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
