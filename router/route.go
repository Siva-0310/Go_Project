package router

import "net/http"

type Route struct {
	path     string
	handlers map[string]*routePath
	group    *string
}
type routePath struct {
	middleware []MiddlewareFunc
	handler    http.HandlerFunc
}
