package rdbms

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/Goboolean/shared/pkg/resolver"
	_ "github.com/lib/pq"
)

type PSQL struct {

	db *sql.DB
	q *Queries

}

var (
	once sync.Once
	instance *PSQL
)

func NewDB(c *resolver.ConfigMap) *PSQL {


	once.Do(func() {
		user, err := c.GetStringKey("USER")
		if err != nil {
			panic(err)
		}


		password, err := c.GetStringKey("PASSWORD")
		if err != nil {
			panic(err)
		}

	

		host, err := c.GetStringKey("HOST")
		if err != nil {
			panic(err)
		}

	

		port, err := c.GetStringKey("PORT")
		if err != nil {
			panic(err)
		}


		database, err := c.GetStringKey("DATABASE")
		if err != nil {
			panic(err)
		}


	
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, database)
	
		db, err := sql.Open("postgres", psqlInfo)
	

		if err != nil {
			panic(err)
		}

		instance = &PSQL{
			db: db,

			q: New(db),

		}
	})

	return instance

}

func (p *PSQL) Close() error {
	return p.db.Close()
}

func (p *PSQL) Ping() error {
	return p.db.Ping()
}

func (p *PSQL) NewQueries() *Queries {
	return p.q
}

func (p *PSQL) NewTx(ctx context.Context) (resolver.Transactioner, error) {
	tx, err := p.db.BeginTx(ctx, nil)
	return NewTransaction(tx, ctx), err
}
