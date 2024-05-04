-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: AddAccountBalance :one
UPDATE accounts
set balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
    RETURNING *;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
OFFSET $1
LIMIT $2
;

-- name: CreateAccount :one
INSERT INTO accounts (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateAccount :one
UPDATE accounts
  set balance = $2
WHERE id = $1
RETURNING *;