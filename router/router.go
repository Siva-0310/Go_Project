package router

import (
	"net/http"
)

type Router struct {
	NotFoundHandler http.HandlerFunc
	Tree            *Node
	middleware      []MiddlewareFunc
}

func (router *Router) addRoute(method string, path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	route := Route{
		method:     method,
		path:       path,
		handler:    handler,
		middleware: mwm,
	}
	arrayPath := []rune(path)
	Insert(arrayPath, router.Tree, &route)
}

func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	arrayPath := []rune(request.URL.Path)
	route := Search(arrayPath, router.Tree)
	if route != nil {
		middleware := append(route.middleware, router.middleware...)
		handler := applyMiddleware(route.handler, middleware...)
		handler(writer, request)
		return
	}
	router.NotFoundHandler(writer, request)
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

// func (router *Router) Include(routerA *Router) {
// 	router.routes = append(router.routes, routerA.routes...)
// }
