package event

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type NewEventInput struct {
	Name             string             `json:"name"`
	Summary          string             `json:"summary"`
	Description      string             `json:"description"`
	Available_ticket int32              `json:"available_ticket"`
	Starting_date    pgtype.Timestamptz `json:"starting_date"`
	Ending_date      pgtype.Timestamptz `json:"ending_date"`
	Venue_id         pgtype.UUID        `json:"venue_id"`
	Host_id          pgtype.UUID        `json:"host_id"`
}
