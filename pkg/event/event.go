package event

import (
	"context"
	"fmt"

	"github.com/gerry-sheva/tixmaster/pkg/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func newEvent(ctx context.Context, dbpool *pgxpool.Pool, p *NewEventInput) error {
	var venue_id pgtype.UUID
	err := venue_id.Scan(p.Venue_id)
	var host_id pgtype.UUID
	err = host_id.Scan(p.Host_id)
	var starting_date pgtype.Timestamptz
	var ending_date pgtype.Timestamptz
	err = starting_date.Scan(p.Starting_date)
	err = ending_date.Scan(p.Ending_date)

	params := sqlc.NewEventParams{
		Name:            p.Name,
		Summary:         p.Summary,
		Description:     p.Description,
		AvailableTicket: p.Available_ticket,
		StartingDate:    starting_date,
		EndingDate:      ending_date,
		VenueID:         venue_id,
		HostID:          host_id,
	}

	fmt.Printf("%v\n", venue_id)

	_, err = sqlc.New(dbpool).NewEvent(ctx, params)
	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}
