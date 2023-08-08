package broker_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Goboolean/shared/pkg/broker"
	"github.com/Goboolean/shared/pkg/resolver"
)

var (
	pub       *broker.Publisher
	data      = &broker.StockAggregate{}
	dataBatch = []*broker.StockAggregate{
		{}, {}, {},
	}
)

func SetupPublisher() {
	var err error

	pub, err = broker.NewPublisher(&resolver.ConfigMap{
		"HOST": os.Getenv("KAFKA_HOST"),
		"PORT": os.Getenv("KAFKA_PORT"),
	})
	if err != nil {
		panic(err)
	}
}

func TeardownPublisher() {
	pub.Close()
}

func TestPublisher(t *testing.T) {

	SetupPublisher()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	if err := pub.Ping(ctx); err != nil {
		t.Errorf("Ping() failed: %v", err)
	}

	TeardownPublisher()
}

func Test_SendData(t *testing.T) {

	var topic = "test-topic"
	SetupPublisher()
	SetupConfigurator()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	exists, err := conf.TopicExists(ctx, topic)
	if err != nil {
		t.Errorf("failed to check topic exists: %v", err)
	}
	if !exists {
		t.Errorf("topic does not exist")
	}

	if err := pub.SendData(topic, data); err != nil {
		t.Errorf("SendData() failed: %v", err)
	}

	TeardownPublisher()
	TeardownConfigurator()
}

func Test_SendDataBatch(t *testing.T) {

	var topic = "test-topic"
	SetupPublisher()
	SetupConfigurator()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	exists, err := conf.TopicExists(ctx, topic)
	if err != nil {
		t.Errorf("failed to check topic exists: %v", err)
	}
	if !exists {
		t.Errorf("topic does not exist")
	}

	if err := pub.SendDataBatch(topic, dataBatch); err != nil {
		t.Errorf("SendDataBatch() failed: %v", err)
	}

	TeardownPublisher()
	TeardownConfigurator()
}
