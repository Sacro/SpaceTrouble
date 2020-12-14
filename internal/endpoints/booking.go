package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/Sacro/SpaceTrouble/internal/ticket"
	"github.com/apex/log"
	"github.com/go-playground/validator/v10"
)

// BookingHandler provides the handler for /bookings
func (h Handler) BookingHandler(w http.ResponseWriter, r *http.Request) {
	var ticket *ticket.Ticket

	err := json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		log.WithError(err).Error("decoding body")
		http.Error(w, "unable to decode body", http.StatusInternalServerError)
		return
	}

	v := validator.New()

	err = v.Struct(ticket)
	if err != nil {
		log.WithError(err).Error("verifying ticket")
		http.Error(w, "unable to verify ticket", http.StatusBadRequest)
		return
	}

	err = h.repository.CreateBooking(ticket)
	if err != nil {
		log.WithError(err).Error("creating booking")
		http.Error(w, "unable to create booking", http.StatusInternalServerError)
		return
	}
}
