package main

import (
	"github.com/gerry-sheva/tixmaster/pkg/api"
	"github.com/gerry-sheva/tixmaster/pkg/database"
	"github.com/gerry-sheva/tixmaster/pkg/util"
)

func main() {
	dbpool := database.ConnectDB()
	defer dbpool.Close()
	rwJSON := util.NewRwJSON()

	api.StartServer(dbpool, rwJSON)
}
