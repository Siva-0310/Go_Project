package router

import "net/http"

type Route struct {
	method     string
	path       string
	handler    http.Handler
	middleware []MiddlewareFunc
}
