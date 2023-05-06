package kafka_test

import (
	"os"
	"testing"

	"github.com/Goboolean/shared-packages/pkg/kafka"
	"github.com/joho/godotenv"
)


type SubscribeListenerImpl struct {}

func (i *SubscribeListenerImpl) OnReceiveMessage(stock *kafka.StockAggregate) {
	received = stock
}

var received *kafka.StockAggregate



func TestMain(m *testing.M) {

	if err := os.Chdir("../../"); err != nil {
		panic(err)
	}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	
	code := m.Run()

	os.Exit(code)
}




func TestConsumer(t *testing.T) {
	sub := kafka.NewConsumer(topic, &SubscribeListenerImpl{})

	if sub == nil {
		t.Errorf("NewConsumer() failed: see log.Fatal")
	}

	pub := kafka.NewPublisher()
	if sub == nil {
		t.Errorf("NewPublisher() failed: see log.Fatal")
	}

	if err := pub.SendData(topic, &kafka.StockAggregate{}); err != nil {
		t.Errorf("NewPublisher() failed: %v", err)
	}

	if received != send {
		t.Errorf("received %v, want %v", received, send)
	}

	if err := sub.Close(); err != nil {
		t.Errorf("NewConsumer() failed: %v", err)
	}
}