-- name: GetEntry :one
SELECT * from entries
where id=$1 LIMIT 1;

-- name: GetListEntry :many
SELECT * from entries
LIMIT $1
OFFSET $2;

-- name: CreateEntry :one
INSERT into entries(account_id,amount)
values
($1,$2) returning *;

-- name: UpdateEntry :one
Update entries
set account_id=$2 , amount=$3
where id=$1 returning *;

-- name: DeleteEntry :exec
Delete from entries
where id=$1;