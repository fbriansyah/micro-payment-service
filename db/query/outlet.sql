-- name: CreateOutlet :one
INSERT INTO outlets ("user", deposit)
VALUES ($1, $2)
RETURNING *;

-- name: GetOutletByUserID :one
SELECT * FROM outlets
WHERE "user" = $1 LIMIT 1;

-- name: UpdateOutletDeposit :one
UPDATE outlets
SET deposit = $1
WHERE id = $2
RETURNING *;