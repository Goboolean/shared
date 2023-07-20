-- name: CreateAccessInfo :exec
INSERT INTO store_log (product_id, "status") VALUES ($1, $2);


-- name: InsertNewStockMeta :exec
INSERT INTO product_meta (id, "name", symbol, "description", "type", exchange, "location") 
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: InsertNewStockPlatformMeta :exec
INSERT INTO product_platform (product_id, platform_name, identifier)
VALUES ($1, $2, $3);


-- name: CheckStockExist :one
SELECT EXISTS(SELECT 1 FROM product_meta WHERE id = ($1));

-- name: GetStockMeta :one
SELECT id, "name", symbol, "description", "type", exchange,  "location"  FROM product_meta WHERE id = ($1);

-- name: GetAllStockMetaList :many
SELECT id, "name", symbol, "description", "type", exchange,  "location"  FROM product_meta;

-- name: GetStockMetaWithPlatform :one
SELECT product_meta.id, "name", symbol, "description", "type", exchange,  "location" , platform_name, identifier 
FROM product_meta 
JOIN product_platform 
ON product_meta.id = product_platform.product_id 
WHERE product_meta.id = ($1);

-- name: UpdatePlatformIdentifier :exec
UPDATE product_platform SET identifier = ($1) WHERE product_id = ($2) AND platform_name = ($3);

-- name: DeletePlatformInfo :exec
DELETE FROM product_platform WHERE product_id = ($1) AND platform_name = ($2);

-- name: InsertPlatformInfo :exec
INSERT INTO product_platform (product_id, platform_name, identifier) VALUES ($1, $2, $3);

-- name: GetStockIdBySymbol :one
SELECT id FROM product_meta WHERE symbol = ($1);