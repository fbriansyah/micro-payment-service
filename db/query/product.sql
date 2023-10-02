-- name: GetProductByCode :one
SELECT * FROM products
WHERE product_code = $1 LIMIT 1;

-- name: CreateProduct :one
INSERT INTO products (product_code, product_name, product_endpoint)
VALUES ($1,$2,$3)
RETURNING *;