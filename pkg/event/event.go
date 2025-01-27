package event

import (
	"context"
	"fmt"
	"os"

	"github.com/gerry-sheva/tixmaster/pkg/database/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/meilisearch/meilisearch-go"
)

func newEvent(ctx context.Context, dbpool *pgxpool.Pool, meilisearch meilisearch.ServiceManager, p *NewEventInput) (sqlc.NewEventRow, error) {
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

	event, err := sqlc.New(dbpool).NewEvent(ctx, params)
	if err != nil {
		return sqlc.NewEventRow{}, err
	}

	index := meilisearch.Index("event")
	task, err := index.AddDocuments(event)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(task.TaskUID)

	return event, nil
}
