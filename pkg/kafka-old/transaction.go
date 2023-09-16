package kafka

import (
	"context"

	"github.com/Goboolean/shared/pkg/resolver"
	"github.com/Shopify/sarama"
)

type Transaction struct {
	session sarama.SyncProducer
	ctx     context.Context
}

func (d *Transaction) Commit() error {
	return d.session.CommitTxn()
}

func (d *Transaction) Rollback() error {
	return d.session.AbortTxn()
}

func (d *Transaction) Context() context.Context {
	return d.ctx
}

func (d *Transaction) Transaction() interface{} {
	return d.session
}

func NewTransaction(session sarama.SyncProducer, ctx context.Context) resolver.Transactioner {
	if err := session.BeginTxn(); err != nil {
		panic(err)
	}
	return &Transaction{session: session, ctx: ctx}
}
