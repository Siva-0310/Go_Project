package router

import "net/http"

type MiddlewareFunc func(http.Handler) http.Handler

func (router *Router) Use(mwm ...MiddlewareFunc) {
	mw := router.middleware
	mw = append(mw, mwm...)
	router.middleware = mw
}

func applyMiddleware(function http.Handler, mwm ...MiddlewareFunc) http.Handler {
	for _, mw := range mwm {
		function = mw(function)
	}
	return function
}
