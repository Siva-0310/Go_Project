package router

import (
	"net/http"
)

func UnprocessableEntityJson(details []byte, writer http.ResponseWriter, reader *http.Request) {
	writer.WriteHeader(http.StatusUnprocessableEntity)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(details)
}

func NotFoundJson(details []byte, writer http.ResponseWriter, reader *http.Request) {
	writer.WriteHeader(http.StatusNotFound)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(details)
}
