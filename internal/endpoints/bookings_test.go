package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sacro/SpaceTrouble/internal/repository"
)

func TestBookingsEndpoint(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/bookings", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	h := NewHandler(&repository.Repository{})
	handler := http.HandlerFunc(h.BookingsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
