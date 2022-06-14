package server

import (
	"context"
	"eleuthHomecast/cast"
	"log"
	"net/http"
)

func Serve(addr string) {
	ctx := context.Background()
	cast := cast.NewHomecast(ctx)
	http.HandleFunc("/", post(HandleTextToSpeeachFunc(ctx, cast)))
	defer func() {
		for _, device := range cast.Devices {
			device.Close()
		}
	}()
	log.Println("Server running...")
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

func post(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodPost)
}

func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			if _, err := writer.Write([]byte("Method Not Allowed")); err != nil {
				log.Println(err)
			}
			return
		}

		writer.Header().Add("Content-Type", "application/json")
		apiFunc(writer, request)
	}
}
