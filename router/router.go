package router

import (
	"encoding/json"
	"net/http"
)

type Route struct {
	path    string
	method  string
	handler func(http.ResponseWriter, *http.Request)
}
type Router struct {
	routes []Route
}

func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for _, route := range router.routes {
		if route.path != request.URL.Path || route.method != request.Method {
			continue
		}
		if err := validator(request); err != "" && request.ContentLength != 0 {
			UnprocessableEntity(err, writer, request)
			return
		}
		route.handler(writer, request)
		return
	}
	NotFound(writer, request)
}

func (router *Router) GET(path string, handler func(http.ResponseWriter, *http.Request)) {
	router.addRoute("GET", path, handler)
}
func (router *Router) POST(path string, handler func(http.ResponseWriter, *http.Request)) {
	router.addRoute("POST", path, handler)
}
func (router *Router) DELETE(path string, handler func(http.ResponseWriter, *http.Request)) {
	router.addRoute("DELETE", path, handler)
}
func (router *Router) PUT(path string, handler func(http.ResponseWriter, *http.Request)) {
	router.addRoute("PUT", path, handler)
}

func (router *Router) IncludeRouter(pathRouter *Router) {
	if pathRouter.routes != nil {
		router.routes = append(router.routes, pathRouter.routes...)
	}
}
func (router *Router) addRoute(method string, path string, handler func(http.ResponseWriter, *http.Request)) {
	temp := Route{
		method:  method,
		path:    path,
		handler: handler,
	}
	router.routes = append(router.routes, temp)
}

func NotFound(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	data := map[string]interface{}{"detail": "no path found"}
	jsonData, err := JSON(data)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusNotFound)
	writer.Write(jsonData)
}

func JSON(data map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func WriteResponce(writer http.ResponseWriter, request *http.Request, data map[string]interface{}, status int) {
	writer.Header().Add("Content-Type", "application/json")
	jsonData, err := JSON(data)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(status)
	writer.Write(jsonData)
}

func Handler(data map[string]interface{}, status int) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		WriteResponce(writer, request, data, status)
	}
}

func validator(request *http.Request) string {
	if request.Header.Get("Content-Type") == "" {
		return "Reques dont have Content-Type header"
	} else if request.Header.Get("Content-Type") != "application/json" {
		return "Request body is not of JSON type"
	}
	return ""
}

func UnprocessableEntity(value string, writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	data := map[string]interface{}{"detail": value}
	jsonData, err := JSON(data)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusUnprocessableEntity)
	writer.Write(jsonData)
}
