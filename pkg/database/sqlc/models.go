// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Event struct {
	EventID         pgtype.UUID
	VenueID         pgtype.UUID
	HostID          pgtype.UUID
	Name            string
	Summary         string
	Description     string
	AvailableTicket int32
	Thumbnail       string
	Banner          string
	StartingDate    pgtype.Timestamptz
	EndingDate      pgtype.Timestamptz
	CreatedAt       pgtype.Timestamptz
	UpdatedAt       pgtype.Timestamptz
	DeletedAt       pgtype.Timestamptz
}

type Host struct {
	HostID    pgtype.UUID
	Name      string
	Avatar    string
	Bio       string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	DeletedAt pgtype.Timestamptz
}

type User struct {
	UserID    pgtype.UUID
	Username  string
	Email     string
	Password  string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type Venue struct {
	VenueID   pgtype.UUID
	Name      string
	Capacity  int32
	City      string
	State     string
	Img       string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	DeletedAt pgtype.Timestamptz
}
