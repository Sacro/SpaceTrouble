package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// BookingHandler provides the handler for /bookings
func (h Handler) BookingHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if id, ok := vars["id"]; ok {
		booking, err := h.repository.Booking(id)
		if err != nil {
			http.Error(w, "booking not found", http.StatusNotFound)
			return
		} else {
			err = json.NewEncoder(w).Encode(booking)
			if err != nil {
				http.Error(w, "Unable to marshal booking", http.StatusInternalServerError)
				return
			}
		}
	}

	http.Error(w, "missing booking id", http.StatusBadRequest)

}
