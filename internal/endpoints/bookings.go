package endpoints

import (
	"encoding/json"
	"net/http"
)

// BookingsHandler provides the handler for /bookings/{id}
func (h Handler) BookingsHandler(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.repository.Bookings()
	if err != nil {
		http.Error(w, "unable to retrieve bookings", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(bookings)

	if err != nil {
		http.Error(w, "unable to marshall bookings", http.StatusInternalServerError)
	}
}
