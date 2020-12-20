package endpoints

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Sacro/SpaceTrouble/internal/spacex"
	"github.com/Sacro/SpaceTrouble/internal/ticket"
	"github.com/apex/log"
	"github.com/go-playground/validator/v10"
)

var ErrDateConflict = errors.New("launch date conflict")

// BookingHandler provides the handler for /bookings
func (h Handler) BookingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var t *ticket.Ticket

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.WithError(err).Error("decoding body")
		http.Error(w, "unable to decode body", http.StatusInternalServerError)
		return
	}

	v := validator.New()

	err = v.Struct(t)
	if err != nil {
		log.WithError(err).Error("verifying ticket")
		http.Error(w, "unable to verify ticket", http.StatusBadRequest)
		return
	}

	launches, err := h.launchClient.GetLaunches(ctx)
	if err != nil {
		log.WithError(err).Error("getting launches")
		http.Error(w, "unable to get launches", http.StatusInternalServerError)
		return
	}

	conflict := checkForLaunchConflicts(t, launches)

	if conflict {
		http.Error(w, "conflicting launch", http.StatusConflict)
		return
	}

	err = h.repository.CreateBooking(t)
	if err != nil {
		log.WithError(err).Error("creating booking")
		http.Error(w, "unable to create booking", http.StatusInternalServerError)
		return
	}
}

func checkForLaunchConflicts(t *ticket.Ticket, launches spacex.LaunchPads) bool {
	for _, launch := range launches {
		// Don't care about past launches
		if !launch.Upcoming {
			continue
		}

		// Check for the same launchpad
		if t.LaunchpadID == string(launch.Launchpad) {
			// Check for conflicting launch date

			if t.LaunchDate.Sub(launch.DateUTC) < time.Hour*24 {
				return true
			}
		}
	}

	return false
}
