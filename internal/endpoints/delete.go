package endpoints

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	err := h.repository.DeleteBooking(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
