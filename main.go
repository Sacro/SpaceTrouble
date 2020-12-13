package main

import (
	"net/http"

	"github.com/Sacro/SpaceTrouble/internal/endpoints"
	"github.com/apex/log"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/booking/{id}", endpoints.BookingHandler)
	r.HandleFunc("/bookings", endpoints.BookingsHandler)

	log.WithError(http.ListenAndServe(":3000", r)).Fatal("http.ListenAndServe")
}
