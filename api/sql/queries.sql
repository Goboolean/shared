-- name: CreateAccessInfo :exec
INSERT INTO store_log (product_id, "status") VALUES (?, ?);

-- name: InsertNewStockMeta :exec
INSERT INTO product_meta (product_id, product_name, symbol, "description", product_type, exchange, "location") 
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: InsertNewStockPlatformMeta :exec
INSERT INTO stock_platform (product_id, platform, identifier) VALUES (?, ?, ?);

-- name: CheckStockExist :one
SELECT EXISTS(SELECT 1 FROM product_meta WHERE product_id = (?));

-- name: GetStockMeta :one
SELECT product_id, product_name, symbol, "description", product_type, exchange,  "location"  FROM product_meta WHERE product_id = (?);

-- name: GetAllStockMetaList :many
SELECT product_id, product_name, symbol, "description", product_type, exchange,  "location"  FROM product_meta;

-- name: GetStockMetaWithPlatform :one
SELECT product_meta.product_id, product_name, symbol, "description", product_type, exchange,  "location" , platform, identifier 
FROM product_meta 
JOIN stock_platform 
ON product_meta.product_id = stock_platform.product_id 
WHERE product_meta.product_id = (?);

-- name: UpdatePlatform:exec
UPDATE stock_platform SET platform = (?), identifier = (?) WHERE product_id = (?);