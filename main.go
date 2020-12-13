package main

import (
	"net/http"

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

	log.WithError(http.ListenAndServe(":3000", r)).Fatal("http.ListenAndServe")
}
