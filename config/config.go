package config

type Config struct {
	MockedKafkaPort string `env:"MOCKED_KAFKA_PORT" envDefault:"9092"`
	AppPort         string `env:"APP_PORT" envDefault:"8080"`
}
