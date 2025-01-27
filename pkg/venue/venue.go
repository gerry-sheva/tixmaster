package venue

import (
	"context"

	"github.com/gerry-sheva/tixmaster/pkg/database/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

func newVenue(ctx context.Context, dbpool *pgxpool.Pool, p NewVenueInput) (sqlc.NewVenueRow, error) {
	params := sqlc.NewVenueParams{
		Name:     p.Name,
		Capacity: p.Capacity,
		City:     p.City,
		State:    p.State,
	}

	venue, err := sqlc.New(dbpool).NewVenue(ctx, params)
	if err != nil {
		return sqlc.NewVenueRow{}, err
	}

	return venue, nil
}
