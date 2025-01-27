package main

import (
	"github.com/gerry-sheva/tixmaster/pkg/api"
	"github.com/gerry-sheva/tixmaster/pkg/database"
	"github.com/meilisearch/meilisearch-go"
)

func main() {
	client := meilisearch.New("http://localhost:7700", meilisearch.WithAPIKey("MASTER_KEY"))

	dbpool := database.ConnectDB()
	defer dbpool.Close()

	api.StartServer(dbpool, client)
}
