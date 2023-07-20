// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: queries.sql

package rdbms

import (
	"context"
	"database/sql"
)

const checkStockExist = `-- name: CheckStockExist :one
SELECT EXISTS(SELECT 1 FROM product_meta WHERE id = ($1))
`

func (q *Queries) CheckStockExist(ctx context.Context, id string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkStockExist, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createAccessInfo = `-- name: CreateAccessInfo :exec
INSERT INTO store_log (product_id, "status") VALUES ($1, $2)
`

type CreateAccessInfoParams struct {
	ProductID string
	Status    string
}

func (q *Queries) CreateAccessInfo(ctx context.Context, arg CreateAccessInfoParams) error {
	_, err := q.db.ExecContext(ctx, createAccessInfo, arg.ProductID, arg.Status)
	return err
}

const deletePlatformInfo = `-- name: DeletePlatformInfo :exec
DELETE FROM product_platform WHERE product_id = ($1) AND platform_name = ($2)
`

type DeletePlatformInfoParams struct {
	ProductID    string
	PlatformName string
}

func (q *Queries) DeletePlatformInfo(ctx context.Context, arg DeletePlatformInfoParams) error {
	_, err := q.db.ExecContext(ctx, deletePlatformInfo, arg.ProductID, arg.PlatformName)
	return err
}

const getAllStockMetaList = `-- name: GetAllStockMetaList :many
SELECT id, "name", symbol, "description", "type", exchange,  "location"  FROM product_meta
`

func (q *Queries) GetAllStockMetaList(ctx context.Context) ([]ProductMetum, error) {
	rows, err := q.db.QueryContext(ctx, getAllStockMetaList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductMetum
	for rows.Next() {
		var i ProductMetum
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Symbol,
			&i.Description,
			&i.Type,
			&i.Exchange,
			&i.Location,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getStockIdBySymbol = `-- name: GetStockIdBySymbol :one
SELECT id FROM product_meta WHERE symbol = ($1)
`

func (q *Queries) GetStockIdBySymbol(ctx context.Context, symbol string) (string, error) {
	row := q.db.QueryRowContext(ctx, getStockIdBySymbol, symbol)
	var id string
	err := row.Scan(&id)
	return id, err
}

const getStockMeta = `-- name: GetStockMeta :one
SELECT id, "name", symbol, "description", "type", exchange,  "location"  FROM product_meta WHERE id = ($1)
`

func (q *Queries) GetStockMeta(ctx context.Context, id string) (ProductMetum, error) {
	row := q.db.QueryRowContext(ctx, getStockMeta, id)
	var i ProductMetum
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Symbol,
		&i.Description,
		&i.Type,
		&i.Exchange,
		&i.Location,
	)
	return i, err
}

const getStockMetaWithPlatform = `-- name: GetStockMetaWithPlatform :one
SELECT product_meta.id, "name", symbol, "description", "type", exchange,  "location" , platform_name, identifier 
FROM product_meta 
JOIN product_platform 
ON product_meta.id = product_platform.product_id 
WHERE product_meta.id = ($1)
`

type GetStockMetaWithPlatformRow struct {
	ID           string
	Name         string
	Symbol       string
	Description  sql.NullString
	Type         string
	Exchange     string
	Location     sql.NullString
	PlatformName string
	Identifier   string
}

func (q *Queries) GetStockMetaWithPlatform(ctx context.Context, id string) (GetStockMetaWithPlatformRow, error) {
	row := q.db.QueryRowContext(ctx, getStockMetaWithPlatform, id)
	var i GetStockMetaWithPlatformRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Symbol,
		&i.Description,
		&i.Type,
		&i.Exchange,
		&i.Location,
		&i.PlatformName,
		&i.Identifier,
	)
	return i, err
}

const insertNewStockMeta = `-- name: InsertNewStockMeta :exec
INSERT INTO product_meta (id, "name", symbol, "description", "type", exchange, "location") 
VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type InsertNewStockMetaParams struct {
	ID          string
	Name        string
	Symbol      string
	Description sql.NullString
	Type        string
	Exchange    string
	Location    sql.NullString
}

func (q *Queries) InsertNewStockMeta(ctx context.Context, arg InsertNewStockMetaParams) error {
	_, err := q.db.ExecContext(ctx, insertNewStockMeta,
		arg.ID,
		arg.Name,
		arg.Symbol,
		arg.Description,
		arg.Type,
		arg.Exchange,
		arg.Location,
	)
	return err
}

const insertNewStockPlatformMeta = `-- name: InsertNewStockPlatformMeta :exec
INSERT INTO product_platform (product_id, platform_name, identifier)
VALUES ($1, $2, $3)
`

type InsertNewStockPlatformMetaParams struct {
	ProductID    string
	PlatformName string
	Identifier   string
}

func (q *Queries) InsertNewStockPlatformMeta(ctx context.Context, arg InsertNewStockPlatformMetaParams) error {
	_, err := q.db.ExecContext(ctx, insertNewStockPlatformMeta, arg.ProductID, arg.PlatformName, arg.Identifier)
	return err
}

const insertPlatformInfo = `-- name: InsertPlatformInfo :exec
INSERT INTO product_platform (product_id, platform_name, identifier) VALUES ($1, $2, $3)
`

type InsertPlatformInfoParams struct {
	ProductID    string
	PlatformName string
	Identifier   string
}

func (q *Queries) InsertPlatformInfo(ctx context.Context, arg InsertPlatformInfoParams) error {
	_, err := q.db.ExecContext(ctx, insertPlatformInfo, arg.ProductID, arg.PlatformName, arg.Identifier)
	return err
}

const updatePlatformIdentifier = `-- name: UpdatePlatformIdentifier :exec
UPDATE product_platform SET identifier = ($1) WHERE product_id = ($2) AND platform_name = ($3)
`

type UpdatePlatformIdentifierParams struct {
	Identifier   string
	ProductID    string
	PlatformName string
}

func (q *Queries) UpdatePlatformIdentifier(ctx context.Context, arg UpdatePlatformIdentifierParams) error {
	_, err := q.db.ExecContext(ctx, updatePlatformIdentifier, arg.Identifier, arg.ProductID, arg.PlatformName)
	return err
}
