package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kamilwoloszyn/kafka-mocked-responser/app"
)

type kafkaRequest struct {
	Message       string
	TypeOfMessage app.MessageType
}

func KafkaHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			var kafkaReqBody kafkaRequest
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&kafkaReqBody); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Header().Set("content-type", "text/plain")
				w.Write(
					[]byte(fmt.Sprintf("Got error : %v", err)),
				)
			}

		}
		w.WriteHeader(http.StatusBadRequest)
	}
}
