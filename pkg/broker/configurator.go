package broker

import (
	"context"
	"fmt"
	"time"

	"github.com/Goboolean/shared/pkg/resolver"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

// Configurator has a role for making and deleting topic, checking topic exists, and getting topic list.
type Configurator struct {
	client *kafka.AdminClient
}

// Constructor throws panic when error occurs
func NewConfigurator(c *resolver.ConfigMap) (*Configurator, error) {

	host, err := c.GetStringKey("HOST")
	if err != nil {
		return nil, err
	}

	port, err := c.GetStringKey("PORT")
	if err != nil {
		return nil, err
	}

	address := fmt.Sprintf("%s:%s", host, port)

	config := &kafka.ConfigMap{
		"bootstrap.servers": address,
		//"debug": "security, broker",
	}

	client, err := kafka.NewAdminClient(config)

	if err != nil {
		return nil, err
	}

	return &Configurator{client: client}, nil
}

// It should be called before program ends to free memory
func (c *Configurator) Close() {

	// If any exception occurs while closing, it will be ignored

	// Error issue:
	// Assertion failed: (r == 0), function rwlock_wrlock, file tinycthread_extra.c, line 157.
	// SIGABRT: abort
	// PC=0x1855ecdb8 m=12 sigcode=0
	// signal arrived during cgo execution

	defer func() {
		recover()
	}()

	c.client.Close()
}

// Check if connection to kafka is alive
func (c *Configurator) Ping(ctx context.Context) error {

	// It requires ctx to be deadline set, otherwise it will return error
	// It will return error if there is no response within deadline
	deadline, ok := ctx.Deadline()

	if !ok {
		return errTimeoutRequired
	}

	remaining := time.Until(deadline)

	_, err := c.client.GetMetadata(nil, true, int(remaining.Milliseconds()))
	return err
}

// Create a topic
func (c *Configurator) CreateTopic(ctx context.Context, topic string) error {

	// It returns error when topic already exists
	topic = packTopic(topic)

	exists, err := c.TopicExists(ctx, topic)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	topicInfo := kafka.TopicSpecification{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}

	result, err := c.client.CreateTopics(ctx, []kafka.TopicSpecification{topicInfo})

	if err != nil {
		return err
	}

	if err := result[0].Error; err.Code() != kafka.ErrNoError {
		return fmt.Errorf(err.String())
	}

	return nil
}

// Delete a topic
func (c *Configurator) DeleteTopic(ctx context.Context, topic string) error {

	// It returns error when the topic does not exist
	topic = packTopic(topic)

	result, err := c.client.DeleteTopics(ctx, []string{topic})

	if err != nil {
		return errors.Wrap(errFatalWhileDeletingTopic, err.Error())
	}

	if err := result[0].Error; err.Code() != kafka.ErrNoError {
		return errors.Wrap(errTrivalWhileDeletingTopic, err.String())
	}

	return nil
}

// Check if given topic exists
func (c *Configurator) TopicExists(ctx context.Context, topic string) (bool, error) {

	topic = packTopic(topic)

	deadline, ok := ctx.Deadline()

	if !ok {
		return false, errTimeoutRequired
	}

	remaining := time.Until(deadline)

	metadata, err := c.client.GetMetadata(nil, true, int(remaining.Milliseconds()))
	if err != nil {
		return false, err
	}

	_, exists := metadata.Topics[topic]
	return exists, nil
}

// Get all existing topic list as a string slice
func (c *Configurator) GetTopicList(ctx context.Context) ([]string, error) {

	deadline, ok := ctx.Deadline()

	if !ok {
		return nil, errTimeoutRequired
	}

	remaining := time.Until(deadline)

	metadata, err := c.client.GetMetadata(nil, true, int(remaining.Milliseconds()))
	if err != nil {
		return nil, err
	}

	topicList := make([]string, 0)

	for topic := range metadata.Topics {
		if len(topic) > 0 {
			topicList = append(topicList, topic)
		}
	}

	return topicList, nil
}
