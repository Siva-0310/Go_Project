package main

import (
	"app/router"
	"fmt"
	"net/http"
)

func NextHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello Next")
}

//	func log(f http.HandlerFunc) http.HandlerFunc {
//		return func(w http.ResponseWriter, r *http.Request) {
//			fmt.Println(r.URL.Path)
//			f(w, r)
//		}
//	}
func main() {
	Router := router.Router{
		NotFoundHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "not")
		}),
		Tree: &router.Node{
			PathLetter: make(map[rune]*router.Node),
			CurrPath:   make([]rune, 0),
		},
	}
	// Router.Use(log)
	Router.Get("/Hello", NextHandler)
	Router.Get("/Hell", NextHandler)
	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: &Router,
	}
	server.ListenAndServe()
}
