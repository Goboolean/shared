package kafka_test

import (
	"context"
	"testing"
	"time"

	"github.com/Goboolean/shared/pkg/kafka"
	"google.golang.org/protobuf/proto"
)

var simEventReceived = make(chan *kafka.SimEvent)

type SimEventTester struct{}

func (t *SimEventTester) OnRecieveSimRequestedEvent(e *kafka.SimEvent) {
	simEventReceived <- e
}

func (t *SimEventTester) OnRecieveSimRequestedRollbackEvent(e *kafka.SimEvent) {
	simEventReceived <- e
}

func (t *SimEventTester) OnRecieveSimPendingEvent(e *kafka.SimEvent) {
	simEventReceived <- e
}

func (t *SimEventTester) OnRecieveSimPendingRollbackEvent(e *kafka.SimEvent) {
	simEventReceived <- e
}

func (t *SimEventTester) OnRecieveSimAllocatedEvent(e *kafka.SimEvent) {
	simEventReceived <- e
}

func (t *SimEventTester) OnRecieveSimAllocatedRollbackEvent(e *kafka.SimEvent) {
	simEventReceived <- e
}

func (t *SimEventTester) OnRecieveSimFailedEvent(e *kafka.SimEvent) {
	simEventReceived <- e
}

func (t *SimEventTester) OnRecieveSimFailedRollbackEvent(e *kafka.SimEvent) {
	simEventReceived <- e
}

func (t *SimEventTester) OnRecieveSimFinishedEvent(e *kafka.SimEvent) {
	simEventReceived <- e
}

func (t *SimEventTester) OnRecieveSimFinishedRollbackEvent(e *kafka.SimEvent) {
	simEventReceived <- e
}

func Test_SimEvent(t *testing.T) {

	type args struct {
		event        *kafka.SimEvent
		subscription func(context.Context)
		sendEvent    func(*kafka.SimEvent) error
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Requested Event",
			args: args{
				event: &kafka.SimEvent{
					Status: int64(kafka.SimEventStatusRequested),
				},
				subscription: func(ctx context.Context) {
					sub.SubscribeSimRequestedEvent(ctx, &SimEventTester{})
				},
				sendEvent: func(event *kafka.SimEvent) error {
					return pub.SendSimRequestedEvent(event)
				},
			},
		},
		{
			name: "Pending Event",
			args: args{
				event: &kafka.SimEvent{
					Status: int64(kafka.SimEventStatusPending),
				},
				subscription: func(ctx context.Context) {
					sub.SubscribeSimPendingEvent(ctx, &SimEventTester{})
				},
				sendEvent: func(event *kafka.SimEvent) error {
					return pub.SendSimPendingEvent(event)
				},
			},
		},
		{
			name: "Allocated Event",
			args: args{
				event: &kafka.SimEvent{
					Status: int64(kafka.SimEventStatusAllocated),
				},
				subscription: func(ctx context.Context) {
					sub.SubscribeSimAllocatedEvent(ctx, &SimEventTester{})
				},
				sendEvent: func(event *kafka.SimEvent) error {
					return pub.SendSimAllocatedEvent(event)
				},
			},
		},
		{
			name: "Failed Event",
			args: args{
				event: &kafka.SimEvent{
					Status: int64(kafka.SimEventStatusFailed),
				},
				subscription: func(ctx context.Context) {
					sub.SubscribeSimFailedEvent(ctx, &SimEventTester{})
				},
				sendEvent: func(event *kafka.SimEvent) error {
					return pub.SendSimFailedEvent(event)
				},
			},
		},
		{
			name: "Finished Event",
			args: args{
				event: &kafka.SimEvent{
					Status: int64(kafka.SimEventStatusFinished),
				},
				subscription: func(ctx context.Context) {
					sub.SubscribeSimFinishedEvent(ctx, &SimEventTester{})
				},
				sendEvent: func(event *kafka.SimEvent) error {
					return pub.SendSimFinishedEvent(event)
				},
			},
		},
	}

	defer func() {
		if err := recover(); err != nil {
			t.Errorf("SubscribeSimEvent() failed: %v", err)
		}
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

			tt.args.subscription(ctx)

			if err := tt.args.sendEvent(tt.args.event); err != nil {
				t.Errorf("SendSimEvent() failed: %v", err)
			}

			select {
			case <-ctx.Done():
				t.Errorf("timeout: failed to receive data")
			case got := <-simEventReceived:
				if !proto.Equal(got, tt.args.event) {
					t.Errorf("OnReceiveSimEvent() = %v, want %v", got, tt.args.event)
				}
			}

			cancel()
		})
	}
}
