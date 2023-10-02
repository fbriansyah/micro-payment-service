-- name: GetProductByCode :one
SELECT * FROM products
WHERE product_code = $1 LIMIT 1;