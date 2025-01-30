package api

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gerry-sheva/tixmaster/pkg/api/middleware"
	"github.com/gerry-sheva/tixmaster/pkg/auth"
	"github.com/gerry-sheva/tixmaster/pkg/event"
	"github.com/gerry-sheva/tixmaster/pkg/host"
	"github.com/gerry-sheva/tixmaster/pkg/venue"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/meilisearch/meilisearch-go"
)

const version = "0.1.0"

type config struct {
	port int
	env  string
}

type app struct {
	cfg         config
	logger      *slog.Logger
	dbpool      *pgxpool.Pool
	meilisearch meilisearch.ServiceManager
	ik          *imagekit.ImageKit
}

func StartServer(dbpool *pgxpool.Pool, meilisearch meilisearch.ServiceManager, ik *imagekit.ImageKit) {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|staging|prod)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &app{
		cfg,
		logger,
		dbpool,
		meilisearch,
		ik,
	}

	usersAPI := auth.New(app.dbpool)
	eventAPI := event.New(app.dbpool, app.meilisearch, ik)
	hostAPI := host.New(app.dbpool, ik)
	venueAPI := venue.New(app.dbpool, app.ik)

	mux := http.NewServeMux()
	mux.Handle("/", middleware.Auth(http.HandlerFunc(app.healthcheckHandler)))
	mux.HandleFunc("POST /register", usersAPI.RegisterUser)
	mux.HandleFunc("POST /login", usersAPI.LoginUser)
	mux.HandleFunc("POST /event", eventAPI.CreateEvent)
	mux.HandleFunc("POST /host", hostAPI.CreateHost)
	mux.HandleFunc("POST /venue", venueAPI.CreateVenue)

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
