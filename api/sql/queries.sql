-- name: CreateAccessInfo :exec
INSERT INTO store_log (product_id, "status") VALUES ($1, $2);

-- name: InsertNewStockMeta :exec
INSERT INTO product_meta (product_id, "name", symbol, "description", "type", exchange, "location") 
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: InsertNewStockPlatformMeta :exec
INSERT INTO stock_platform (product_id, platform, identifier) VALUES ($1, $2, $3);

-- name: CheckStockExist :one
SELECT EXISTS(SELECT 1 FROM product_meta WHERE product_id = ($1));

-- name: GetStockMeta :one
SELECT product_id, "name", symbol, "description", "type", exchange,  "location"  FROM product_meta WHERE product_id = ($1);

-- name: GetAllStockMetaList :many
SELECT product_id, "name", symbol, "description", "type", exchange,  "location"  FROM product_meta;

-- name: GetStockMetaWithPlatform :one
SELECT product_meta.product_id, "name", symbol, "description", "type", exchange,  "location" , platform, identifier 
FROM product_meta 
JOIN stock_platform 
ON product_meta.product_id = stock_platform.product_id 
WHERE product_meta.product_id = ($1);

-- name: UpdatePlatform :exec
UPDATE stock_platform SET platform = ($1), identifier = ($2) WHERE product_id = ($3);