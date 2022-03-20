package web

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// map k--> url and v --> handler
type Engine struct {
	router map[string]HandlerFunc
}

// return router map
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// add route
func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
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
	key := r.Method + "-" + r.URL.Path
	// 判断路由的处理 handler 是否存在
	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

// Run http server
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
