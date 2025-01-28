package auth

import (
	"context"
	"errors"

	sqlc "github.com/gerry-sheva/tixmaster/pkg/database/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

func register(ctx context.Context, dbpool *pgxpool.Pool, p *AuthInput) (string, error) {
	password_hash, err := hashPassword(p.Password)
	if err != nil {
		return "", err
	}

	params := sqlc.NewUserParams{
		Email:    p.Email,
		Username: p.Username,
		Password: password_hash,
	}

	user, err := sqlc.New(dbpool).NewUser(ctx, params)
	if err != nil {
		return "", err
	}

	jwt, err := createJWT(user)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func login(ctx context.Context, dbpool *pgxpool.Pool, p *AuthInput) (string, error) {
	params := sqlc.GetUserParams{
		Email:    p.Email,
		Username: p.Username,
	}

	user, err := sqlc.New(dbpool).GetUser(ctx, params)
	if err != nil {
		return "", err
	}

	match, _, err := verifyPassword(p.Password, user.Password)
	if err != nil {
		return "", err
	}

	if !match {
		println("HWHWH")
		return "", errors.New("Invalid credentials")
	}

	jwt, err := createJWT(user.Email)
	if err != nil {
		return "", err
	}

	return jwt, nil
}
