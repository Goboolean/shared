package kafka

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Goboolean/common/pkg/resolver"
	"github.com/Shopify/sarama"
)

type Producer struct {
	producer sarama.SyncProducer
}

func NewProducer(c *resolver.ConfigMap) (*Producer, error) {

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
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Transaction.ID = createTransactionID()

	producer, err := sarama.NewSyncProducer([]string{address}, config)
	if err != nil {
		return nil, err
	}

	return &Producer{producer: producer}, nil
}

func createTransactionID() string {
	pid := os.Getpid()
	pidString := strconv.Itoa(pid)
	return pidString
}

func (p *Producer) Close() error {
	return p.producer.Close()
}
