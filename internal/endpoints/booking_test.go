package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sacro/SpaceTrouble/internal/repository"
)

func TestBookingEndpoint(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/booking/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	h := NewHandler(&repository.Repository{})
	handler := http.HandlerFunc(h.BookingHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
