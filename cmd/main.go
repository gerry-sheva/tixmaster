package main

import (
	"github.com/gerry-sheva/tixmaster/pkg/api"
	"github.com/gerry-sheva/tixmaster/pkg/database"
)

func main() {
	dbpool := database.ConnectDB()
	defer dbpool.Close()

	api.StartServer(dbpool)
}
