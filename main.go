package main

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/kamilwoloszyn/kafka-mocked-responser/config"
)

func main() {
	cfg := config.Config{}
	l := log.Default()
	if err := env.Parse(&cfg); err != nil {
		l.Fatalf("unable to parse config: %v", err)
	}

}
