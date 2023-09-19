package main

import (
	"app/router"
	"fmt"
	"net/http"
)

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	data := map[string]interface{}{"data": "HelloWorld"}
	router.WriteResponce(writer, request, data, http.StatusAccepted)
}

func NextHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello Next")
}

func main() {
	Router := router.Router{}
	Router.GET("/", router.Handler(map[string]interface{}{"data": "HelloWorld"}, http.StatusAccepted))
	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: &Router,
	}
	server.ListenAndServe()
}
