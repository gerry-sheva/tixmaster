-- name: GetUser :one
select user_id, username, email from users where user_id=$1;

-- name: NewUser :one
insert into users (username, email, password) values ($1, $2, $3) returning user_id, username, email;
