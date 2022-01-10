package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kamilwoloszyn/kafka-mocked-responser/http/handlers"
)

func Init() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("v1", func(r chi.Router) {
		r.Post("/kafka", handlers.KafkaHandler())
	})

}
