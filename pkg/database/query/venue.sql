-- name: NewVenue :one
insert into venue (name, capacity, city, state)
values ($1, $2, $3, $4)
returning venue_id, name, capacity, city, state;
