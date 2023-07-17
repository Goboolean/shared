package kafka

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
	"google.golang.org/protobuf/proto"
)



func (p *Producer) sendRealEvent(status RealEventStatus, event *RealEvent) error {

	event.Status = int64(status)

	data, err := proto.Marshal(event)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: RealEventTopic[RealEventStatus(event.Status)],
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = p.producer.SendMessage(msg);
	return err
}

func (p *Producer) sendRealRollbackEvent(status RealEventStatus, event *RealEvent) error {

	event.Status = int64(status)

	data, err := proto.Marshal(event)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: RealEventRollbackTopic[RealEventStatus(event.Status)],
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = p.producer.SendMessage(msg);
	return err
}



func (p *Producer) SendRealRequestedEvent(event *RealEvent) error {
	return p.sendRealEvent(RealEventStatusRequested, event)
}

func (p *Producer) SendRealRequestedRollbackEvent(event *RealEvent) error {
	return p.sendRealRollbackEvent(RealEventStatusRequested, event)
}

func (p *Producer) SendRealPendingEvent(event *RealEvent) error {
	return p.sendRealEvent(RealEventStatusPending, event)
}

func (p *Producer) SendRealPendingRollbackEvent(event *RealEvent) error {
	return p.sendRealRollbackEvent(RealEventStatusPending, event)
}

func (p *Producer) SendRealAllocatedEvent(event *RealEvent) error {
	return p.sendRealEvent(RealEventStatusAllocated, event)
}

func (p *Producer) SendRealAllocatedRollbackEvent(event *RealEvent) error {
	return p.sendRealRollbackEvent(RealEventStatusAllocated, event)
}

func (p *Producer) SendRealCeasedEvent(event *RealEvent) error {
	return p.sendRealEvent(RealEventStatusCeased, event)
}

func (p *Producer) SendRealCeasedRollbackEvent(event *RealEvent) error {
	return p.sendRealRollbackEvent(RealEventStatusCeased, event)
}

func (p *Producer) SendRealFailedEvent(event *RealEvent) error {
	return p.sendRealEvent(RealEventStatusFailed, event)
}

func (p *Producer) SendRealFailedRollbackEvent(event *RealEvent) error {
	return p.sendRealRollbackEvent(RealEventStatusFailed, event)
}



func (c *Consumer) subscribeRealEvent(ctx context.Context, topic string, callback func(*RealEvent)) error {
	
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

				var event *RealEvent	
				proto.Unmarshal(message.Value, event)

				callback(event)
			}
		}
	}(ctx)

	return nil
}



type RealRequestedEventListener interface {
	OnRecieveRealRequestedEvent(*RealEvent)
	OnRecieveRealRequestedRollbackEvent(*RealEvent)
}

func (c *Consumer) SubscribeRealRequestedEvent(ctx context.Context, impl RealRequestedEventListener) {
	if err := c.subscribeRealEvent(ctx, RealEventTopic[RealEventStatusRequested], func(event *RealEvent) {
		impl.OnRecieveRealRequestedEvent(event)
	}); err != nil {
		panic(err)
	}

	if err := c.subscribeRealEvent(ctx, RealEventRollbackTopic[RealEventStatusRequested], func(event *RealEvent) {
		impl.OnRecieveRealRequestedRollbackEvent(event)
	}); err != nil {
		panic(err)
	}
}



type RealPendingEventListener interface {
	OnRecieveRealPendingEvent(*RealEvent)
	OnRecieveRealPendingRollbackEvent(*RealEvent)
}

func (c *Consumer) SubscribeRealPendingEvent(ctx context.Context, impl RealPendingEventListener) {
	if err := c.subscribeRealEvent(ctx, RealEventTopic[RealEventStatusRequested], func(event *RealEvent) {
		impl.OnRecieveRealPendingEvent(event)
	}); err != nil {
		panic(err)
	}

	if err := c.subscribeRealEvent(ctx, RealEventRollbackTopic[RealEventStatusRequested], func(event *RealEvent) {
		impl.OnRecieveRealPendingRollbackEvent(event)
	}); err != nil {
		panic(err)
	}
}



type RealAllocatedEventListener interface {
	OnRecieveRealAllocatedEvent(*RealEvent)
	OnRecieveRealAllocatedRollbackEvent(*RealEvent)
}

func (c *Consumer) SubscribeRealAllocatedEvent(ctx context.Context, impl RealAllocatedEventListener) {
	if err := c.subscribeRealEvent(ctx, RealEventTopic[RealEventStatusRequested], func(event *RealEvent) {
		impl.OnRecieveRealAllocatedEvent(event)
	}); err != nil {
		panic(err)
	}

	if err := c.subscribeRealEvent(ctx, RealEventRollbackTopic[RealEventStatusRequested], func(event *RealEvent) {
		impl.OnRecieveRealAllocatedRollbackEvent(event)
	}); err != nil {
		panic(err)
	}
}



type RealCeasedEventListener interface {
	OnRecieveRealCeasedEvent(*RealEvent)
	OnRecieveRealCeasedRollbackEvent(*RealEvent)
}

func (c *Consumer) SubscribeRealCeasedEvent(ctx context.Context, impl RealCeasedEventListener) {
	if err := c.subscribeRealEvent(ctx, RealEventTopic[RealEventStatusRequested], func(event *RealEvent) {
		impl.OnRecieveRealCeasedEvent(event)
	}); err != nil {
		panic(err)
	}

	if err := c.subscribeRealEvent(ctx, RealEventRollbackTopic[RealEventStatusRequested], func(event *RealEvent) {
		impl.OnRecieveRealCeasedRollbackEvent(event)
	}); err != nil {
		panic(err)
	}
}



type RealFailedEventListener interface {
	OnRecieveRealFailedEvent(*RealEvent)
	OnRecieveRealFailedRollbackEvent(*RealEvent)
}

func (c *Consumer) SubscribeRealFailedEvent(ctx context.Context, impl RealFailedEventListener) {
	if err := c.subscribeRealEvent(ctx, RealEventTopic[RealEventStatusRequested], func(event *RealEvent) {
		impl.OnRecieveRealFailedEvent(event)
	}); err != nil {
		panic(err)
	}

	if err := c.subscribeRealEvent(ctx, RealEventRollbackTopic[RealEventStatusRequested], func(event *RealEvent) {
		impl.OnRecieveRealFailedRollbackEvent(event)
	}); err != nil {
		panic(err)
	}
}