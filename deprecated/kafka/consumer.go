package kafka

import (
	"fmt"

	"github.com/Goboolean/common/pkg/resolver"
	"github.com/IBM/sarama"
)

type Consumer struct {
	consumer sarama.Consumer

	data map[string]chan interface{}
}

func NewConsumer(c *resolver.ConfigMap) (*Consumer, error) {

	host, err := c.GetStringKey("HOST")
	if err != nil {
		return nil, err
	}

	port, err := c.GetStringKey("PORT")
	if err != nil {
		return nil, err
	}

	address := fmt.Sprintf("%s:%s", host, port)

	config := sarama.NewConfig()
	config.Producer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{address}, config)

	if err != nil {
		return nil, err
	}

	return &Consumer{
		consumer: consumer,
	}, nil
}

func (c *Consumer) Close() error {
	return c.consumer.Close()
}
