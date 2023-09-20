package main

import (
	"app/router"
	"fmt"
	"net/http"
)

func NextHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello Next")
}
func log(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		f.ServeHTTP(w, r)
	})
}
func main() {
	Router := router.Router{
		NotFoundHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "not")
		}),
	}
	Router.Use(log)
	Router.Get("/", NextHandler)
	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: &Router,
	}
	server.ListenAndServe()
}
