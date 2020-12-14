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

	launches, err := h.getLaunches()
	if err != nil {
		log.WithError(err).Error("getting launches")
		http.Error(w, "unable to get launches", http.StatusInternalServerError)
		return
	}

	for _, launch := range launches {
		// Don't care about past launches
		if !launch.Upcoming {
			continue
		}

		// Check for the same launchpad
		if ticket.LaunchpadID == string(launch.Launchpad) {
			// Check for conflicting launch date
			if ticket.LaunchDate == launch.DateUTC {
				http.Error(w, "conflicting launch", http.StatusConflict)
				return
			}
		}

		continue
	}

	err = h.repository.CreateBooking(ticket)
	if err != nil {
		log.WithError(err).Error("creating booking")
		http.Error(w, "unable to create booking", http.StatusInternalServerError)
		return
	}
}
