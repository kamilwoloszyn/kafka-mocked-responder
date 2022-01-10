package app

type MessageType string

var (
	JsonType  MessageType = "json"
	PlainText MessageType = "plain"
)

type Kafka struct {
}

func (k *Kafka) Response(message string, t MessageType) {

}
