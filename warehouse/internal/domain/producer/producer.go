package producer

type Publisher interface {
	Produce(topic string, payload interface{}) error
}
