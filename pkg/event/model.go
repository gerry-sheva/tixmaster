package event

import "time"

type NewEventInput struct {
	Name             string    `json:"name"`
	Summary          string    `json:"summary"`
	Description      string    `json:"description"`
	Available_ticket int32     `json:"available_ticket"`
	Starting_date    time.Time `json:"starting_date"`
	Ending_date      time.Time `json:"ending_date"`
	Venue_id         string    `json:"venue_id"`
	Host_id          string    `json:"host_id"`
}
