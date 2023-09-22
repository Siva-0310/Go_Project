package router

import (
	"fmt"
	"net/http"
)

type router struct {
	NotFoundHandler http.HandlerFunc
	Tree            *Node
	middleware      []MiddlewareFunc
	group           map[string][]MiddlewareFunc
}

func App(NotFoundHandler http.HandlerFunc) router {
	return router{
		NotFoundHandler: NotFoundHandler,
		middleware:      make([]MiddlewareFunc, 0),
		group:           make(map[string][]MiddlewareFunc),
		Tree: &Node{
			PathLetter: make(map[rune]*Node),
			CurrPath:   make([]rune, 0),
		},
	}
}

func (router *router) addRoute(method string, path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	val := Search([]rune(path), router.Tree)
	if val != nil {
		val.handlers[method] = &routePath{
			middleware: mwm,
			handler:    handler,
		}
	} else {
		route := Route{
			path: path,
			handlers: map[string]*routePath{method: {
				middleware: mwm,
				handler:    handler,
			}},
			group: nil,
		}
		arrayPath := []rune(path)
		Insert(arrayPath, router.Tree, &route)
	}
}

func (router *router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	arrayPath := []rune(request.URL.Path)
	route := Search(arrayPath, router.Tree)
	if route != nil {
		var middleware []MiddlewareFunc
		if route.group != nil {
			middleware = append(router.group[*route.group], router.middleware...)
			middleware = append(route.handlers[request.Method].middleware, middleware...)
		} else {
			middleware = append(route.handlers[request.Method].middleware, router.middleware...)
		}

		handler := applyMiddleware(route.handlers[request.Method].handler, middleware...)
		handler(writer, request)
		return
	}
	router.NotFoundHandler(writer, request)
}

func (rtr *router) Get(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	rtr.addRoute("GET", path, handler, mwm...)
}
func (rtr *router) Post(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	rtr.addRoute("POST", path, handler, mwm...)
}
func (rtr *router) Delete(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	rtr.addRoute("DELETE", path, handler, mwm...)
}
func (rtr *router) Put(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	rtr.addRoute("PUT", path, handler, mwm...)
}
func (rtr *router) Patch(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	rtr.addRoute("PATCH", path, handler, mwm...)
}
func (rtr *router) Head(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	rtr.addRoute("HEAD", path, handler, mwm...)
}
func (rtr *router) Options(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	rtr.addRoute("OPTIONS", path, handler, mwm...)
}
func (rtr *router) Trace(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	rtr.addRoute("TRACE", path, handler, mwm...)
}
func (rtr *router) Connect(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	rtr.addRoute("CONNECT", path, handler, mwm...)
}

func (rtr *router) Include(groups ...group) {
	for _, group := range groups {
		for _, route := range group.routes {
			fmt.Println(route.handlers)
			Insert([]rune(route.path), rtr.Tree, &route)
		}
		rtr.group[group.prefix] = group.middleware
	}
}
