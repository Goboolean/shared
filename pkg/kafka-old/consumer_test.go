package kafka_test

import (
	"os"
	"testing"

	"github.com/Goboolean/shared/pkg/kafka"
	"github.com/Goboolean/shared/pkg/resolver"
)

var sub *kafka.Consumer

func SetupConsumer() {
	var err error

	sub, err = kafka.NewConsumer(&resolver.ConfigMap{
		"HOST": os.Getenv("KAFKA_HOST"),
		"PORT": os.Getenv("KAFKA_PORT"),
	})
	if err != nil {
		panic(err)
	}
}

func TeardownConsumer() {
	if err := sub.Close(); err != nil {
		panic(err)
	}
}

func Test_Consumer(t *testing.T) {
	SetupConsumer()
	TeardownConsumer()
}