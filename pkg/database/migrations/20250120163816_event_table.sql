-- +goose Up
-- +goose StatementBegin
create table venue (
    venue_id uuid primary key default uuid_generate_v1mc (),
    name text collate "case_insensitive" unique not null,
    capacity integer not null,
    city text not null,
    state text not null,
    created_at timestamptz not null default now (),
    updated_at timestamptz,
    deleted_at timestamptz
);

select
    trigger_updated_at ('venue');

create table host (
    host_id uuid primary key default uuid_generate_v1mc (),
    name text collate "case_insensitive" unique not null,
    avatar text not null,
    bio text not null,
    created_at timestamptz not null default now (),
    updated_at timestamptz,
    deleted_at timestamptz
);

select
    trigger_updated_at ('host');

create table event (
    event_id uuid primary key default uuid_generate_v1mc (),
    venue_id uuid not null,
    host_id uuid not null,
    name text collate "case_insensitive" unique not null,
    summary text not null,
    description text not null,
    available_ticket integer not null,
    starting_date timestamptz not null,
    ending_date timestamptz not null,
    created_at timestamptz not null default now (),
    updated_at timestamptz,
    deleted_at timestamptz,
    foreign key (venue_id) references venue (venue_id),
    foreign key (host_id) references host (host_id)
);

select
    trigger_updated_at ('event');

-- +goose StatementEnd
