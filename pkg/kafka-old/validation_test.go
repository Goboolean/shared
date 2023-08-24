package kafka_test

import (
	"context"
	"testing"
	"time"

	"github.com/Goboolean/shared/pkg/kafka"
	"google.golang.org/protobuf/proto"
)

var valEventReceived = make(chan *kafka.SimEvent)

type ValEventTester struct{}

func (t *ValEventTester) OnRecieveValRequestedEvent(e *kafka.SimEvent) {
	valEventReceived <- e
}

func (t *ValEventTester) OnRecieveValRequestedRollbackEvent(e *kafka.SimEvent) {
	valEventReceived <- e
}

func (t *ValEventTester) OnRecieveValPendingEvent(e *kafka.SimEvent) {
	valEventReceived <- e
}

func (t *ValEventTester) OnRecieveValPendingRollbackEvent(e *kafka.SimEvent) {
	valEventReceived <- e
}

func (t *ValEventTester) OnRecieveValAllocatedEvent(e *kafka.SimEvent) {
	valEventReceived <- e
}

func (t *ValEventTester) OnRecieveValAllocatedRollbackEvent(e *kafka.SimEvent) {
	valEventReceived <- e
}

func (t *ValEventTester) OnRecieveValFailedEvent(e *kafka.SimEvent) {
	valEventReceived <- e
}

func (t *ValEventTester) OnRecieveValFailedRollbackEvent(e *kafka.SimEvent) {
	valEventReceived <- e
}

func (t *ValEventTester) OnRecieveValFinishedEvent(e *kafka.SimEvent) {
	valEventReceived <- e
}

func (t *ValEventTester) OnRecieveValFinishedRollbackEvent(e *kafka.SimEvent) {
	valEventReceived <- e
}

func Test_ValEvent(t *testing.T) {

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
					Status: int64(kafka.ValEventStatusRequested),
				},
				subscription: func(ctx context.Context) {
					sub.SubscribeValRequestedEvent(ctx, &ValEventTester{})
				},
				sendEvent: func(event *kafka.SimEvent) error {
					return pub.SendValRequestedEvent(event)
				},
			},
		},
		{
			name: "Pending Event",
			args: args{
				event: &kafka.SimEvent{
					Status: int64(kafka.ValEventStatusPending),
				},
				subscription: func(ctx context.Context) {
					sub.SubscribeValPendingEvent(ctx, &ValEventTester{})
				},
				sendEvent: func(event *kafka.SimEvent) error {
					return pub.SendValPendingEvent(event)
				},
			},
		},
		{
			name: "Allocated Event",
			args: args{
				event: &kafka.SimEvent{
					Status: int64(kafka.ValEventStatusAllocated),
				},
				subscription: func(ctx context.Context) {
					sub.SubscribeValAllocatedEvent(ctx, &ValEventTester{})
				},
				sendEvent: func(event *kafka.SimEvent) error {
					return pub.SendValAllocatedEvent(event)
				},
			},
		},
		{
			name: "Failed Event",
			args: args{
				event: &kafka.SimEvent{
					Status: int64(kafka.ValEventStatusFailed),
				},
				subscription: func(ctx context.Context) {
					sub.SubscribeValFailedEvent(ctx, &ValEventTester{})
				},
				sendEvent: func(event *kafka.SimEvent) error {
					return pub.SendValFailedEvent(event)
				},
			},
		},
		{
			name: "Finished Event",
			args: args{
				event: &kafka.SimEvent{
					Status: int64(kafka.ValEventStatusFinished),
				},
				subscription: func(ctx context.Context) {
					sub.SubscribeValFinishedEvent(ctx, &ValEventTester{})
				},
				sendEvent: func(event *kafka.SimEvent) error {
					return pub.SendValFinishedEvent(event)
				},
			},
		},
	}

	defer func() {
		if err := recover(); err != nil {
			t.Errorf("SubscribeValEvent() failed: %v", err)
		}
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

			tt.args.subscription(ctx)

			if err := tt.args.sendEvent(tt.args.event); err != nil {
				t.Errorf("SendValEvent() failed: %v", err)
			}

			select {
			case <-ctx.Done():
				t.Errorf("timeout: failed to receive data")
			case got := <-valEventReceived:
				if !proto.Equal(got, tt.args.event) {
					t.Errorf("OnReceiveValEvent() = %v, want %v", got, tt.args.event)
				}
			}

			cancel()
		})
	}
}
