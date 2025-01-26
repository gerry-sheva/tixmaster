package event

import (
	"context"

	"github.com/gerry-sheva/tixmaster/pkg/database/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

func newEvent(ctx context.Context, dbpool *pgxpool.Pool, p *NewEventInput) error {
	params := sqlc.NewEventParams{
		Name:            p.Name,
		Summary:         p.Summary,
		Description:     p.Description,
		AvailableTicket: p.Available_ticket,
		StartingDate:    p.Starting_date,
		EndingDate:      p.Ending_date,
		VenueID:         p.Venue_id,
		HostID:          p.Host_id,
	}

	_, err := sqlc.New(dbpool).NewEvent(ctx, params)
	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}
