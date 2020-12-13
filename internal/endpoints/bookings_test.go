package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBookingsEndpoint(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/bookings", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BookingsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
