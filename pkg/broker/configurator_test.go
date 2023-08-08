package broker_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Goboolean/shared/pkg/broker"
	"github.com/Goboolean/shared/pkg/resolver"
)

var conf *broker.Configurator

func SetupConfigurator() {

	var err error
	conf, err = broker.NewConfigurator(&resolver.ConfigMap{
		"HOST": os.Getenv("KAFKA_HOST"),
		"PORT": os.Getenv("KAFKA_PORT"),
	})
	if err != nil {
		panic(err)
	}
}

func TeardownConfigurator() {
	conf.Close()
}

func Test_Configurator(t *testing.T) {
	SetupConfigurator()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	if err := conf.Ping(ctx); err != nil {
		t.Errorf("Ping() = %v", err)
	}

	cancel()
	TeardownConfigurator()
}

func Test_CreateTopic(t *testing.T) {

	var topic = "test-topic"
	SetupConfigurator()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	if err := conf.CreateTopic(ctx, topic); err != nil {
		t.Errorf("CreateTopic() = %v", err)
	}

	exists, err := conf.TopicExists(ctx, topic)
	if err != nil {
		t.Errorf("TopicExists() = %v", err)
	}
	if !exists {
		t.Errorf("TopicExists() = %v, expected = true", exists)
	}

	TeardownConfigurator()
}

func Test_DeleteTopic(t *testing.T) {

	var topic = "test-topic"
	SetupConfigurator()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	if err := conf.DeleteTopic(ctx, topic); err != nil {
		t.Errorf("DeleteTopic() = %v", err)
	}

	exists, err := conf.TopicExists(ctx, topic)
	if err != nil {
		t.Errorf("TopicExists() = %v", err)
	}
	if exists {
		t.Errorf("TopicExists() = %v, expected = false", exists)
	}

	TeardownConfigurator()
}

func Test_GetTopicList(t *testing.T) {

	SetupConfigurator()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	topicList, err := conf.GetTopicList(ctx)
	if err != nil {
		t.Errorf("GetTopicList() = %v", err)
	}

	fmt.Printf("Topic Count: %d\n", len(topicList))
	fmt.Printf("Topic List: \n")
	for _, topic := range topicList {
		fmt.Println(topic)
	}

	TeardownConfigurator()
}
