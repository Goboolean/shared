package kafka_test

import (
	"context"
	"testing"
	"time"

	"github.com/Goboolean/shared/pkg/kafka"
	"google.golang.org/protobuf/proto"
)

var realEventReceived = make(chan *kafka.RealEvent)

type RealEventTester struct{}

func (t *RealEventTester) OnRecieveRealRequestedEvent(e *kafka.RealEvent) {
	realEventReceived <- e
}

func (t *RealEventTester) OnRecieveRealRequestedRollbackEvent(e *kafka.RealEvent) {
	realEventReceived <- e
}

func (t *RealEventTester) OnRecieveRealPendingEvent(e *kafka.RealEvent) {
	realEventReceived <- e
}

func (t *RealEventTester) OnRecieveRealPendingRollbackEvent(e *kafka.RealEvent) {
	realEventReceived <- e
}

func (t *RealEventTester) OnRecieveRealAllocatedEvent(e *kafka.RealEvent) {
	realEventReceived <- e
}

func (t *RealEventTester) OnRecieveRealAllocatedRollbackEvent(e *kafka.RealEvent) {
	realEventReceived <- e
}

func Test_RealEvent(t *testing.T) {

	type args struct {
		event        *kafka.RealEvent
		subscription func(context.Context)
		sendEvent    func(*kafka.RealEvent) error
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Requested Event",
			args: args{
				event: &kafka.RealEvent{
					Status: int64(kafka.RealEventStatusRequested),
				},
				subscription: func(ctx context.Context) {
					sub.SubscribeRealRequestedEvent(ctx, &RealEventTester{})
				},
				sendEvent: func(event *kafka.RealEvent) error {
					return pub.SendRealRequestedEvent(event)
				},
			},
		},
		{
			name: "Pending Event",
			args: args{
				event: &kafka.RealEvent{
					Status: int64(kafka.RealEventStatusPending),
				},
				subscription: func(ctx context.Context) {
					sub.SubscribeRealPendingEvent(ctx, &RealEventTester{})
				},
				sendEvent: func(event *kafka.RealEvent) error {
					return pub.SendRealPendingEvent(event)
				},
			},
		},
		{
			name: "Allocated Event",
			args: args{
				event: &kafka.RealEvent{
					Status: int64(kafka.RealEventStatusAllocated),
				},
				subscription: func(ctx context.Context) {
					sub.SubscribeRealAllocatedEvent(ctx, &RealEventTester{})
				},
				sendEvent: func(event *kafka.RealEvent) error {
					return pub.SendRealAllocatedEvent(event)
				},
			},
		},
	}

	defer func() {
		if err := recover(); err != nil {
			t.Errorf("SubscribeRealEvent() failed: %v", err)
		}
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

			tt.args.subscription(ctx)

			if err := tt.args.sendEvent(tt.args.event); err != nil {
				t.Errorf("SendRealEvent() failed: %v", err)
			}

			select {
			case <-ctx.Done():
				t.Errorf("timeout: failed to receive data")
			case got := <-realEventReceived:
				if !proto.Equal(got, tt.args.event) {
					t.Errorf("OnReceiveRealEvent() = %v, want %v", got, tt.args.event)
				}
			}

			cancel()
		})
	}
}
