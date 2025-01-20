package auth

import (
	"context"

	tixmaster "github.com/gerry-sheva/tixmaster/pkg/database/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

func register(ctx context.Context, dbpool *pgxpool.Pool, p *AuthInput) error {
	password_hash, err := hashPassword(p.Password)
	if err != nil {
		return err
	}

	params := tixmaster.NewUserParams{
		Email:    p.Email,
		Username: p.Username,
		Password: password_hash,
	}

	_, err = tixmaster.New(dbpool).NewUser(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
