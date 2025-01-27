-- name: NewVenue :one
insert into venue (name, capacity, city, state, img)
values ($1, $2, $3, $4, $5)
returning venue_id, name, capacity, city, state, img;
