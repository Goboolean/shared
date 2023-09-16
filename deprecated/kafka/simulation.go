package kafka

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
	"google.golang.org/protobuf/proto"
)



func (p *Producer) sendSimEvent(status SimEventStatus, event *SimEvent) error {

	event.Status = int64(status)

	data, err := proto.Marshal(event)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: SimEventTopic[SimEventStatus(event.Status)],
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = p.producer.SendMessage(msg);
	return err
}

func (p *Producer) sendSimRollbackEvent(status SimEventStatus, event *SimEvent) error {

	event.Status = int64(status)

	data, err := proto.Marshal(event)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: SimEventRollbackTopic[SimEventStatus(event.Status)],
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = p.producer.SendMessage(msg);
	return err
}




func (p *Producer) SendSimRequestedEvent(event *SimEvent) error {
	return p.sendSimEvent(SimEventStatusRequested, event)
}

func (p *Producer) SendSimRequestedRollbackEvent(event *SimEvent) error {
	return p.sendSimRollbackEvent(SimEventStatusRequested, event)
}

func (p *Producer) SendSimPendingEvent(event *SimEvent) error {
	return p.sendSimEvent(SimEventStatusPending, event)
}

func (p *Producer) SendSimPendingRollbackEvent(event *SimEvent) error {
	return p.sendSimRollbackEvent(SimEventStatusPending, event)
}

func (p *Producer) SendSimAllocatedEvent(event *SimEvent) error {
	return p.sendSimEvent(SimEventStatusAllocated, event)
}

func (p *Producer) SendSimAllocatedRollbackEvent(event *SimEvent) error {
	return p.sendSimRollbackEvent(SimEventStatusAllocated, event)
}

func (p *Producer) SendSimFailedEvent(event *SimEvent) error {
	return p.sendSimEvent(SimEventStatusFailed, event)
}

func (p *Producer) SendSimFailedRollbackEvent(event *SimEvent) error {
	return p.sendSimRollbackEvent(SimEventStatusFailed, event)
}

func (p *Producer) SendSimFinishedEvent(event *SimEvent) error {
	return p.sendSimEvent(SimEventStatusFinished, event)
}

func (p *Producer) SendSimFinishedRollbackEvent(event *SimEvent) error {
	return p.sendSimRollbackEvent(SimEventStatusFinished, event)
}



func (c *Consumer) subscribeSimEvent(ctx context.Context, topic string, callback func(*SimEvent)) error {
	
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



type SimRequestedEventListener interface {
	OnRecieveSimRequestedEvent(*SimEvent)
	OnRecieveSimRequestedRollbackEvent(*SimEvent)
}

func (c *Consumer) SubscribeSimRequestedEvent(ctx context.Context, impl SimRequestedEventListener) {
	if err := c.subscribeSimEvent(ctx, SimEventTopic[SimEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveSimRequestedEvent(event)
	}); err != nil {
		panic(err)
	}

	if err := c.subscribeSimEvent(ctx, SimEventRollbackTopic[SimEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveSimRequestedRollbackEvent(event)
	}); err != nil {
		panic(err)
	}
}



type SimPendingEventListener interface {
	OnRecieveSimPendingEvent(*SimEvent)
	OnRecieveSimPendingRollbackEvent(*SimEvent)
}

func (c *Consumer) SubscribeSimPendingEvent(ctx context.Context, impl SimPendingEventListener) {
	if err := c.subscribeSimEvent(ctx, SimEventTopic[SimEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveSimPendingEvent(event)
	}); err != nil {
		panic(err)
	}

	if err := c.subscribeSimEvent(ctx, SimEventRollbackTopic[SimEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveSimPendingRollbackEvent(event)
	}); err != nil {
		panic(err)
	}
}


type SimAllocatedEventListener interface {
	OnRecieveSimAllocatedEvent(*SimEvent)
	OnRecieveSimAllocatedRollbackEvent(*SimEvent)
}

func (c *Consumer) SubscribeSimAllocatedEvent(ctx context.Context, impl SimAllocatedEventListener) {
	if err := c.subscribeSimEvent(ctx, SimEventTopic[SimEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveSimAllocatedEvent(event)
	}); err != nil {
		panic(err)
	}

	if err := c.subscribeSimEvent(ctx, SimEventRollbackTopic[SimEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveSimAllocatedRollbackEvent(event)
	}); err != nil {
		panic(err)
	}
}


type SimFailedEventListener interface {
	OnRecieveSimFailedEvent(*SimEvent)
	OnRecieveSimFailedRollbackEvent(*SimEvent)
}

func (c *Consumer) SubscribeSimFailedEvent(ctx context.Context, impl SimFailedEventListener) {
	if err := c.subscribeSimEvent(ctx, SimEventTopic[SimEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveSimFailedEvent(event)
	}); err != nil {
		panic(err)
	}

	if err := c.subscribeSimEvent(ctx, SimEventRollbackTopic[SimEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveSimFailedRollbackEvent(event)
	}); err != nil {
		panic(err)
	}
}


type SimFinishedEventListener interface {
	OnRecieveSimFinishedEvent(*SimEvent)
	OnRecieveSimFinishedRollbackEvent(*SimEvent)
}

func (c *Consumer) SubscribeSimFinishedEvent(ctx context.Context, impl SimFinishedEventListener) {
	if err := c.subscribeSimEvent(ctx, SimEventTopic[SimEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveSimFinishedEvent(event)
	}); err != nil {
		panic(err)
	}

	if err := c.subscribeSimEvent(ctx, SimEventRollbackTopic[SimEventStatusRequested], func(event *SimEvent) {
		impl.OnRecieveSimFinishedRollbackEvent(event)
	}); err != nil {
		panic(err)
	}
}