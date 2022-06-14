package server

import (
	"context"
	"eleuthHomecast/cast"
	"eleuthHomecast/domain"
	"encoding/json"
	"log"
	"net/http"
)

//text to speeach
func HandleTextToSpeeachFunc(ctx context.Context, cast *cast.NewDevice) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody domain.TtsRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			BadRequest(w, "Bad Request")
			return
		}
		//defaultLang := flag.String("lang", "en", "Default language to speak")
		//flag.Parse()
		if requestBody.Text == "" {
			log.Printf("[INFO] Skip request due to no text given")
			BadRequest(w, "Bad Request")
			return
		}
		if requestBody.Lang == "" {
			requestBody.Lang = "en"
		}
		for _, device := range cast.Devices {
			if err := device.Speak(ctx, requestBody.Text, requestBody.Lang); err != nil {
				log.Printf("[ERROR] Failed to speak: %v", err)
			}
		}
	}
}
