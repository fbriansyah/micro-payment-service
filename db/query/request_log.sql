-- name: CreateRequestLog :one
INSERT INTO request_logs 
(mode, product, bill_number, name, total_amount, biller_response, outlet)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetRequestLogByID :one
SELECT * FROM request_logs
WHERE
    id = $1
LIMIT 1;