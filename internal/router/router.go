// Route path, route mapping
package router

import (
	"rawhttp/internal/request"
	"rawhttp/internal/response"
)

type HandlerFunc func(req *request.Request) *response.Response

type Router struct {
	routes map[string]map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]HandlerFunc),
	}
}
func (r *Router) GET(path string, handler HandlerFunc) {
	if r.routes["GET"] == nil {
		r.routes["GET"] = make(map[string]HandlerFunc)
	}
	r.routes["GET"][path] = handler
}

func (r *Router) POST(path string, handler HandlerFunc) {
	if r.routes["POST"] == nil {
		r.routes["POST"] = make(map[string]HandlerFunc)
	}
	r.routes["POST"][path] = handler
}

func (r *Router) Handle(req *request.Request) *response.Response {
	if handlers, ok := r.routes[req.Method]; ok {
		if handler, ok := handlers[req.Path]; ok {
			return handler(req)
		}
	}
	return &response.Response{
		StatusCode: 404,
		Body:       []byte("Not Found"),
	}
}
