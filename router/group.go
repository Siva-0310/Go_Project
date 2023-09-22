package router

import (
	"net/http"
)

type group struct {
	routes     []Route
	prefix     string
	middleware []MiddlewareFunc
}

func (group *group) addRoute(method string, path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	path = group.prefix + path
	var temp *Route
	for _, route := range group.routes {
		if route.path == path {
			temp = &route
			break
		}
	}
	if temp == nil {
		route := Route{
			path: path,
			handlers: map[string]*routePath{method: {
				middleware: mwm,
				handler:    handler,
			}},
			group: &group.prefix,
		}
		group.routes = append(group.routes, route)
	} else {
		temp.handlers[method] = &routePath{
			middleware: mwm,
			handler:    handler,
		}
	}
}
func Group(prefix string) group {
	return group{
		prefix:     prefix,
		middleware: make([]MiddlewareFunc, 0),
		routes:     make([]Route, 0),
	}
}
func (grp *group) Get(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	grp.addRoute("GET", path, handler, mwm...)
}
func (grp *group) Post(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	grp.addRoute("POST", path, handler, mwm...)
}
func (grp *group) Delete(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	grp.addRoute("DELETE", path, handler, mwm...)
}
func (grp *group) Put(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	grp.addRoute("PUT", path, handler, mwm...)
}
func (grp *group) Patch(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	grp.addRoute("PATCH", path, handler, mwm...)
}
func (grp *group) Head(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	grp.addRoute("HEAD", path, handler, mwm...)
}
func (grp *group) Options(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	grp.addRoute("OPTIONS", path, handler, mwm...)
}
func (grp *group) Trace(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	grp.addRoute("TRACE", path, handler, mwm...)
}
func (grp *group) Connect(path string, handler http.HandlerFunc, mwm ...MiddlewareFunc) {
	grp.addRoute("CONNECT", path, handler, mwm...)
}
