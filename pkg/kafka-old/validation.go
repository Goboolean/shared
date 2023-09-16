package kafka

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
	"google.golang.org/protobuf/proto"
)



func (p *Producer) sendValEvent(status ValEventStatus, event *SimEvent) error {

	event.Status = int64(status)

	data, err := proto.Marshal(event)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: ValEventTopic[ValEventStatus(event.Status)],
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = p.producer.SendMessage(msg);
	return err
}

func (p *Producer) sendValRollbackEvent(status ValEventStatus, event *SimEvent) error {

	event.Status = int64(status)

	data, err := proto.Marshal(event)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: ValEventRollbackTopic[ValEventStatus(event.Status)],
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = p.producer.SendMessage(msg);
	return err
}




func (p *Producer) SendValRequestedEvent(event *SimEvent) error {
	return p.sendValEvent(ValEventStatusRequested, event)
}

func (p *Producer) SendValRequestedRollbackEvent(event *SimEvent) error {
	return p.sendValRollbackEvent(ValEventStatusRequested, event)
}

func (p *Producer) SendValPendingEvent(event *SimEvent) error {
	return p.sendValEvent(ValEventStatusPending, event)
}

func (p *Producer) SendValPendingRollbackEvent(event *SimEvent) error {
	return p.sendValRollbackEvent(ValEventStatusPending, event)
}

func (p *Producer) SendValAllocatedEvent(event *SimEvent) error {
	return p.sendValEvent(ValEventStatusAllocated, event)
}

func (p *Producer) SendValAllocatedRollbackEvent(event *SimEvent) error {
	return p.sendValRollbackEvent(ValEventStatusAllocated, event)
}

func (p *Producer) SendValFailedEvent(event *SimEvent) error {
	return p.sendValEvent(ValEventStatusFailed, event)
}

func (p *Producer) SendValFailedRollbackEvent(event *SimEvent) error {
	return p.sendValRollbackEvent(ValEventStatusFailed, event)
}

func (p *Producer) SendValFinishedEvent(event *SimEvent) error {
	return p.sendValEvent(ValEventStatusFinished, event)
}

func (p *Producer) SendValFinishedRollbackEvent(event *SimEvent) error {
	return p.sendValRollbackEvent(ValEventStatusFinished, event)
}



func (c *Consumer) subscribeValEvent(ctx context.Context, topic string, callback func(*SimEvent)) error {
	
	pc, err := c.consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		return err
	}

	go func(ctx context.Context) {

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			for err := range pc.Errors() {
				log.Fatal(err)
			}

			for message := range pc.Messages() {

				var event *SimEvent	
				proto.Unmarshal(message.Value, event)

				callback(event)
			}
		}
	}(ctx)

	return nil
}



type ValRequestedEventListener interface {
	OnRecieveValRequestedEvent(*SimEvent)
	OnRecieveValRequestedRollbackEvent(*SimEvent)
}

func (c *Consumer) SubscribeValRequestedEvent(ctx context.Context, impl ValRequestedEventListener) {
	if err := c.subscribeValEvent(ctx, ValEventTopic[ValEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveValRequestedEvent(event)
	}); err != nil {
		panic(err)
	}

	if err := c.subscribeValEvent(ctx, ValEventRollbackTopic[ValEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveValRequestedRollbackEvent(event)
	}); err != nil {
		panic(err)
	}
}



type ValPendingEventListener interface {
	OnRecieveValPendingEvent(*SimEvent)
	OnRecieveValPendingRollbackEvent(*SimEvent)
}

func (c *Consumer) SubscribeValPendingEvent(ctx context.Context, impl ValPendingEventListener) {
	if err := c.subscribeValEvent(ctx, ValEventTopic[ValEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveValPendingEvent(event)
	}); err != nil {
		panic(err)
	}

	if err := c.subscribeValEvent(ctx, ValEventRollbackTopic[ValEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveValPendingRollbackEvent(event)
	}); err != nil {
		panic(err)
	}
}


type ValAllocatedEventListener interface {
	OnRecieveValAllocatedEvent(*SimEvent)
	OnRecieveValAllocatedRollbackEvent(*SimEvent)
}

func (c *Consumer) SubscribeValAllocatedEvent(ctx context.Context, impl ValAllocatedEventListener) {
	if err := c.subscribeValEvent(ctx, ValEventTopic[ValEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveValAllocatedEvent(event)
	}); err != nil {
		panic(err)
	}

	if err := c.subscribeValEvent(ctx, ValEventRollbackTopic[ValEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveValAllocatedRollbackEvent(event)
	}); err != nil {
		panic(err)
	}
}


type ValFailedEventListener interface {
	OnRecieveValFailedEvent(*SimEvent)
	OnRecieveValFailedRollbackEvent(*SimEvent)
}

func (c *Consumer) SubscribeValFailedEvent(ctx context.Context, impl ValFailedEventListener) {
	if err := c.subscribeValEvent(ctx, ValEventTopic[ValEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveValFailedEvent(event)
	}); err != nil {
		panic(err)
	}

	if err := c.subscribeValEvent(ctx, ValEventRollbackTopic[ValEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveValFailedRollbackEvent(event)
	}); err != nil {
		panic(err)
	}
}


type ValFinishedEventListener interface {
	OnRecieveValFinishedEvent(*SimEvent)
	OnRecieveValFinishedRollbackEvent(*SimEvent)
}

func (c *Consumer) SubscribeValFinishedEvent(ctx context.Context, impl ValFinishedEventListener) {
	if err := c.subscribeValEvent(ctx, ValEventTopic[ValEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveValFinishedEvent(event)
	}); err != nil {
		panic(err)
	}

	if err := c.subscribeValEvent(ctx, ValEventRollbackTopic[ValEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveValFinishedRollbackEvent(event)
	}); err != nil {
		panic(err)
	}
}