package adapter

import "github.com/kamilwoloszyn/kafka-mocked-responser/app"

type Responser interface {
	Response(message string, t app.MessageType)
}
