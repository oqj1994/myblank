-- name: GetTransfer :one
SELECT * from transfers
where id=$1 
LIMIT 1;

-- name: GetListTransfer :many
SELECT * from transfers 
LIMIT $1
OFFSET $2;

-- name: CreateTransfer :one
INSERT into transfers
(from_account_id,to_account_id,amount)
values
($1,$2,$3)
returning *;

-- name: UpdateTransfer :one
Update transfers 
set from_account_id=$2,
to_account_id=$3,
amount=$4
where id=$1
returning *;

-- name: DeleteTransfer :exec
Delete from transfers
where id=$1;
