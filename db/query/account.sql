
-- name: CreateAccount :one
INSERT INTO accounts (
                      owner,
                      balance,
                      currency
) VALUES (
    $1,$2,$3 /* THE INPUT ARGUMENT for the insert statement */
) RETURNING *; /* return all the values of the inserted row */


-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
order by id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;