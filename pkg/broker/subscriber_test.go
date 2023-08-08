package broker_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Goboolean/shared/pkg/broker"
	"github.com/Goboolean/shared/pkg/resolver"
	"google.golang.org/protobuf/proto"
)

var (
	sub *broker.Subscriber
)

func SetupSubscriber() {
	var err error

	sub, err = broker.NewSubscriber(&resolver.ConfigMap{
		"HOST": os.Getenv("KAFKA_HOST"),
		"PORT": os.Getenv("KAFKA_PORT"),
	}, context.Background(), &SubscribeListenerImpl{})
	if err != nil {
		panic(err)
	}
}

func TeardownSubscriber() {
	sub.Close()
}

func Test_Subscriber(t *testing.T) {

	SetupSubscriber()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	if err := sub.Ping(ctx); err != nil {
		t.Errorf("Ping() failed: %v", err)
	}

	TeardownSubscriber()
}

type SubscribeListenerImpl struct{}

var stockChan = make(chan *broker.StockAggregate)

func (i *SubscribeListenerImpl) OnReceiveStockAggs(name string, data *broker.StockAggregate) {
	stockChan <- data
}

func Test_Subscribe(t *testing.T) {

	var topic = "test-topic"
	SetupSubscriber()
	SetupPublisher()
	SetupConfigurator()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	exists, err := conf.TopicExists(ctx, topic)
	if err != nil {
		t.Errorf("failed to check topic exists: %v", err)
	}
	if !exists {
		t.Errorf("topic does not exist")
	}

	type args struct {
		topic string
		data  *broker.StockAggregate
	}

	tests := []struct {
		name string
		args args
		want *broker.StockAggregate
	}{
		{
			name: "send mock data",
			args: args{
				topic: topic,
				data: &broker.StockAggregate{
					Average: 1234,
				},
			},
			want: &broker.StockAggregate{
				Average: 1234,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

			if err := sub.Subscribe(topic); err != nil {
				t.Errorf("Subscribe() = %v", err)
			}

			if err := pub.SendData(tt.args.topic, tt.args.data); err != nil {
				t.Errorf("SendData() = %v", err)
			}

			select {
			case <-ctx.Done():
				t.Errorf("timeout: failed to receive data")
			case got := <-stockChan:
				if !proto.Equal(got, tt.want) {
					t.Errorf("OnReceiveStockAggs() = %v, want %v", got, tt.want)
				}
			}

			cancel()
		})
	}

	if err := pub.SendData(topic, data); err != nil {
		t.Errorf("NewPublisher() failed: %v", err)
	}

	close(stockChan)

	TeardownSubscriber()
	TeardownPublisher()
	TeardownConfigurator()
}
