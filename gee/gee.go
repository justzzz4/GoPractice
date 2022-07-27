package gee

import (
	"net/http"
)

type HandleFunc func(*Context)

type Engine struct {
	r *router
}

func New() *Engine {
	return &Engine{
		r: newRouter(),
	}
}

func (e *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	e.r.addRoute(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandleFunc) {
	e.r.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandleFunc) {
	e.r.addRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.r.handle(c)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
