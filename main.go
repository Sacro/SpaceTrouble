package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Sacro/SpaceTrouble/internal/endpoints"
	"github.com/Sacro/SpaceTrouble/internal/repository"
	"github.com/apex/log"
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
)

func main() {

	db := pg.Connect(&pg.Options{})
	repo := repository.New(db)
	handler := endpoints.NewHandler(repo)

	r := mux.NewRouter()
	r.HandleFunc("/booking/{id}", handler.BookingHandler)
	r.HandleFunc("/bookings", handler.BookingsHandler)

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
		if err := srv.ListenAndServe(); err != nil {
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	log.Info("shutting down")
	os.Exit(0)
}
