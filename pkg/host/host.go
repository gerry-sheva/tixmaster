package host

import (
	"context"

	"github.com/gerry-sheva/tixmaster/pkg/database/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

func newHost(ctx context.Context, dbpool *pgxpool.Pool, p *NewHostInput) error {
	params := sqlc.NewHostParams{
		Name:   p.Name,
		Avatar: p.Avatar,
		Bio:    p.Bio,
	}

	_, err := sqlc.New(dbpool).NewHost(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
