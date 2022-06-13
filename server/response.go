package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func Success(writer http.ResponseWriter, response interface{}) {
	if response == nil {
		return
	}
	data, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		InternalServerError(writer, "marshal error")
		return
	}
	if _, err := writer.Write(data); err != nil {
		log.Println(err)
	}
}

func BadRequest(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusBadRequest, message)
}

func InternalServerError(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusInternalServerError, message)
}

func httpError(writer http.ResponseWriter, code int, message string) {
	data, _ := json.Marshal(errorResponse{
		Code:    code,
		Message: message,
	})
	writer.WriteHeader(code)
	if data != nil {
		if _, err := writer.Write(data); err != nil {
			log.Println(err)
			return
		}
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
