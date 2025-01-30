package testhelper

import (
	"context"
	"path/filepath"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/meilisearch"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	ConnectionString string
}

type MeilisearchContainer struct {
	*meilisearch.MeilisearchContainer
	Host string
	Port string
}

func CreateTestContainer(ctx context.Context) (*PostgresContainer, error) {
	pgContainer, err := postgres.Run(ctx, "postgres",
		postgres.WithInitScripts(filepath.Join("..", "database", "test", "migrations.sql.gz")),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		PostgresContainer: pgContainer,
		ConnectionString:  connStr,
	}, nil
}

func CreateMeilisearchContainer(ctx context.Context) (*MeilisearchContainer, error) {
	meiliContainer, err := meilisearch.Run(
		ctx,
		"getmeili/meilisearch:v1.12.8",
		meilisearch.WithMasterKey("MASTER_KEY"),
	)

	if err != nil {
		return nil, err
	}

	host, err := meiliContainer.Host(ctx)
	if err != nil {
		return nil, err
	}
	port, err := meiliContainer.MappedPort(ctx, "7700")
	if err != nil {
		return nil, err
	}

	return &MeilisearchContainer{
		Host:                 host,
		Port:                 port.Port(),
		MeilisearchContainer: meiliContainer,
	}, nil
}
