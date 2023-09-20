package router

import (
	"net/http"
)

type Router struct {
	NotFoundHandler http.Handler
	routes          []Route
	middleware      []MiddlewareFunc
}

func (router *Router) addRoute(method string, path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	route := Route{
		method:     method,
		path:       path,
		handler:    handler,
		middleware: mwm,
	}
	router.routes = append(router.routes, route)
}

func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for _, route := range router.routes {
		if route.method == request.Method && route.path == request.URL.Path {
			middleware := append(route.middleware, router.middleware...)
			handler := applyMiddleware(route.handler, middleware...)
			handler.ServeHTTP(writer, request)
			return
		}
	}
	router.NotFoundHandler.ServeHTTP(writer, request)
}

func (router *Router) Get(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	router.addRoute("GET", path, handler, mwm...)
}
func (router *Router) Post(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	router.addRoute("POST", path, handler, mwm...)
}
func (router *Router) Delete(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	router.addRoute("DELETE", path, handler, mwm...)
}
func (router *Router) Put(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	router.addRoute("PUT", path, handler, mwm...)
}
func (router *Router) Include(routerA *Router) {
	router.routes = append(router.routes, routerA.routes...)
}
