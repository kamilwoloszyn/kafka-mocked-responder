package service

import "github.com/kamilwoloszyn/kafka-mocked-responser/adapter"

type Service struct {
	HttpService adapter.HTTPMaker
	Kafka       adapter.Responser
}

func New(
	httpService adapter.HTTPMaker,
	kafka adapter.Responser,
) *Service {
	return &Service{
		HttpService: httpService,
		Kafka:       kafka,
	}
}
