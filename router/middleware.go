package router

import "net/http"

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func (router *router) Use(mwm ...MiddlewareFunc) {
	mw := router.middleware
	mw = append(mw, mwm...)
	router.middleware = mw
}
func (grp *group) Use(mwm ...MiddlewareFunc) {
	mw := grp.middleware
	mw = append(mw, mwm...)
	grp.middleware = mw
}
func applyMiddleware(function http.HandlerFunc, mwm ...MiddlewareFunc) http.HandlerFunc {
	for _, mw := range mwm {
		function = mw(function)
	}
	return function
}
