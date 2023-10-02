-- name: CreateTransaction :one
INSERT INTO transactions
(
    bill_number, 
    product, 
    inquiry, 
    payment, 
    amount, 
    refference_number, 
    outlet, 
    status
)
VALUES
($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;

-- name: GetTransactionByID :one
SELECT * FROM transactions WHERE id = $1 LIMIT 1;