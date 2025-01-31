package main

import (
	"log"
	"os"

	"github.com/gerry-sheva/tixmaster/pkg/api"
	"github.com/gerry-sheva/tixmaster/pkg/database"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/joho/godotenv"
	"github.com/meilisearch/meilisearch-go"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client := meilisearch.New("http://localhost:7700", meilisearch.WithAPIKey("MASTER_KEY"))
	ik, err := imagekit.New()
	if err != nil {
		panic(err)
	}

	dbpool := database.ConnectDB(os.Getenv("DATABASE_URL"))
	defer dbpool.Close()

	print("Starting server")
	api.StartServer(dbpool, client, ik)
}
