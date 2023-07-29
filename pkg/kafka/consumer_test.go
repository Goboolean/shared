package kafka_test

import (
	"os"
	"testing"

	"github.com/Goboolean/shared/pkg/kafka"
	"github.com/Goboolean/shared/pkg/resolver"
)

var sub *kafka.Consumer

func SetupConsumer() {
	sub = kafka.NewConsumer(&resolver.ConfigMap{
		"HOST": os.Getenv("KAFKA_HOST"),
		"PORT": os.Getenv("KAFKA_PORT"),
	})
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