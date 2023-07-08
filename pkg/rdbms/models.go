// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package rdbms

import (
	"database/sql"
	"time"
)

type StockMetum struct {
	Hash        string
	StockName   string
	Symbol      string
	Description sql.NullString
	ProductType string
	Exchange    string
	Location    string
}

type StockPlatform struct {
	Platform   string
	Identifier string
	StockHash  string
}

type StoreLog struct {
	StockHash string
	StoredAt  time.Time
	Status    string
}
