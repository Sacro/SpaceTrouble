package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/Sacro/SpaceTrouble/internal/endpoints"
	"github.com/Sacro/SpaceTrouble/internal/repository"
	"github.com/Sacro/SpaceTrouble/internal/spacex"
	"github.com/apex/log"
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/peterbourgon/ff/v3"
)

func main() {
	appName := filepath.Base(os.Args[0])
	fs := flag.NewFlagSet(appName, flag.ExitOnError)

	var (
		username = fs.String("postgres-user", "postgres", "postgres username")
		password = fs.String("postgres-password", "postgres", "postgres password")
		hostname = fs.String("postgres-hostname", "localhost", "postgres hostname")
		database = fs.String("postgres-database", "spacetrouble", "database name")
	)

	err := ff.Parse(fs, os.Args[1:], ff.WithEnvVarNoPrefix())
	if err != nil {
		log.WithError(err).Fatal("parsing arguments")
	}

	db := pg.Connect(&pg.Options{
		Addr:     *hostname,
		User:     *username,
		Password: *password,
		Database: *database,
	})

	client := &http.Client{
		Timeout: time.Second * 15,
	}

	repo := repository.New(db)
	err = repo.CreateSchema()
	if err != nil {
		log.WithError(err).Fatalf("migrating database")
	}

	launchClient := &spacex.HTTPLaunchClient{}

	handler := endpoints.NewHandler(client, launchClient, repo)

	r := mux.NewRouter()
	r.HandleFunc("/bookings", handler.BookingHandler).Methods("POST")
	r.HandleFunc("/bookings", handler.BookingsHandler).Methods("GET")

	srv := http.Server{
		Addr: "0.0.0.0:3000",

		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err = srv.ListenAndServe(); err != nil {
			log.WithError(err).Fatal("srv.ListenAndServe")
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	timeout := time.Second * 15
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err = srv.Shutdown(ctx)
	if err != nil {
		log.WithError(err).Fatal("shutting down server")
	}

	log.Info("shutting down")

	go func() {
		os.Exit(0)
	}()
}
