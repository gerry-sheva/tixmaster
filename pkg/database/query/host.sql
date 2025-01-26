-- name: NewHost :one
insert into host (name, avatar, bio) values ($1, $2, $3) returning host_id, name, avatar, bio;
