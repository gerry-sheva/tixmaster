-- name: NewEvent :one
with inserted_event as (
    insert into event (name, summary, description, available_ticket, starting_date, ending_date, venue_id, host_id)
    values ($1, $2, $3, $4, $5, $6, $7, $8)
    returning event_id, name, summary, starting_date, ending_date, venue_id, host_id
)
select
    ie.event_id as id,
    ie.name,
    ie.summary,
    ie.starting_date,
    ie.ending_date,
    v.name as venue_name,
    v.city,
    v.state,
    h.name as host_name,
    h.avatar as host_avatar,
    h.bio as host_bio
from
    inserted_event ie
join
    venue v on ie.venue_id = v.venue_id
join
    host h on ie.host_id = h.host_id;
