-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
for NO KEY Update;

-- name: GetListAccount :many
SELECT * FROM accounts 
LIMIT $1
OFFSET $2;

-- name: CreateAccount :one
INSERT into accounts(owner,balance,currency)
values ($1,$2,$3) returning *;

-- name: DeleteAccount :exec
Delete  from accounts 
where id= $1 ;

-- name: UpdateAccount :one
Update accounts
set balance = $2
where id=$1 returning *; 

-- name: AddAccountBalance :one
Update accounts
set balance = balance+ sqlc.arg(amount)
where id=$1 returning *; 