package main

import (
	"app/router"
	"fmt"
	"net/http"
)

func NextHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello Next")
}
func log(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		f(w, r)
	}
}

func main() {
	Router := router.App(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "NotFound")
	})
	group := router.Group("/hello")
	group.Get("/", NextHandler)
	Router.Get("/", NextHandler)
	for i := 0; i <= 10000; i += 1 {
		Router.Get(fmt.Sprintf("/%d", i), NextHandler)
	}
	for i := 0; i <= 3; i += 1 {
		group.Get(fmt.Sprintf("/%d", i), NextHandler)
	}
	//group.Use(log)
	Router.Include(group)
	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: &Router,
	}
	server.ListenAndServe()
}
