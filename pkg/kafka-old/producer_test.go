package kafka_test

import (
	"os"
	"testing"

	"github.com/Goboolean/shared/pkg/kafka-old"
	"github.com/Goboolean/shared/pkg/resolver"
)

var pub *kafka.Producer

func SetupProducer() {
	var err error
	pub, err = kafka.NewProducer(&resolver.ConfigMap{
		"HOST": os.Getenv("KAFKA_HOST"),
		"PORT": os.Getenv("KAFKA_PORT"),
	})
	if err != nil {
		panic(err)
	}
}

func TeardownProducer() {
	if err := pub.Close(); err != nil {
		panic(err)
	}
}

func TestProducer(t *testing.T) {
	SetupProducer()
	TeardownProducer()
}
