package auth

import (
	"context"

	sqlc "github.com/gerry-sheva/tixmaster/pkg/database/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

func register(ctx context.Context, dbpool *pgxpool.Pool, p *AuthInput) error {
	password_hash, err := hashPassword(p.Password)
	if err != nil {
		return err
	}

	params := sqlc.NewUserParams{
		Email:    p.Email,
		Username: p.Username,
		Password: password_hash,
	}

	_, err = sqlc.New(dbpool).NewUser(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
