package api

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gerry-sheva/tixmaster/pkg/auth"
	"github.com/gerry-sheva/tixmaster/pkg/event"
	"github.com/gerry-sheva/tixmaster/pkg/host"
	"github.com/gerry-sheva/tixmaster/pkg/util"
	"github.com/gerry-sheva/tixmaster/pkg/venue"
	"github.com/jackc/pgx/v5/pgxpool"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type app struct {
	cfg    config
	logger *slog.Logger
	dbpool *pgxpool.Pool
	rwJSON *util.RwJSON
}

func StartServer(dbpool *pgxpool.Pool, rwJSON *util.RwJSON) {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|staging|prod)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &app{
		cfg,
		logger,
		dbpool,
		rwJSON,
	}

	usersAPI := auth.New(app.dbpool, app.rwJSON)
	eventAPI := event.New(app.dbpool, app.rwJSON)
	hostAPI := host.New(app.dbpool, app.rwJSON)
	venueAPI := venue.New(app.dbpool, app.rwJSON)

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.healthcheckHandler)
	mux.HandleFunc("/register", usersAPI.RegisterUser)
	mux.HandleFunc("POST /event", eventAPI.NewEvent)
	mux.HandleFunc("POST /host", hostAPI.NewHost)
	mux.HandleFunc("POST /venue", venueAPI.NewVenue)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
