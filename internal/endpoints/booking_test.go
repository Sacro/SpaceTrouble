package endpoints

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sacro/SpaceTrouble/internal/repository"
	"github.com/Sacro/SpaceTrouble/internal/spacex"
	"github.com/Sacro/SpaceTrouble/internal/ticket"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBookingEndpointValidData(t *testing.T) {
	f := faker.New()
	fakeTicket := ticket.Fake(f)

	body, err := json.Marshal(fakeTicket)
	assert.Nil(t, err)

	req, err := http.NewRequestWithContext(context.Background(), "POST", "/bookings", bytes.NewBuffer(body))

	assert.Nil(t, err)

	rr := httptest.NewRecorder()

	db := &repository.MockRepository{}
	db.On("CreateBooking", mock.Anything).Return(nil)

	lc := &spacex.MockLaunchClient{}

	h := NewHandler(&http.Client{}, lc, db)
	handler := http.HandlerFunc(h.BookingHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	db.AssertExpectations(t)
}

func TestBookingEndpointInvalidData(t *testing.T) {
	fakeTicket := ticket.Ticket{}

	body, err := json.Marshal(fakeTicket)
	assert.Nil(t, err)

	req, err := http.NewRequestWithContext(context.Background(), "POST", "/bookings", bytes.NewBuffer(body))

	assert.Nil(t, err)

	rr := httptest.NewRecorder()

	db := &repository.MockRepository{}
	lc := &spacex.MockLaunchClient{}

	h := NewHandler(&http.Client{}, lc, db)
	handler := http.HandlerFunc(h.BookingHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	db.AssertExpectations(t)
}

func TestBookingEndpointConflictingData(t *testing.T) {
	f := faker.New()
	fakeTicket := ticket.Fake(f)

	body, err := json.Marshal(fakeTicket)
	assert.Nil(t, err)

	req, err := http.NewRequestWithContext(context.Background(), "POST", "/bookings", bytes.NewBuffer(body))

	assert.Nil(t, err)

	rr := httptest.NewRecorder()

	db := &repository.MockRepository{}
	db.On("CreateBooking", mock.Anything).Return(nil)

	lc := &spacex.MockLaunchClient{}

	h := NewHandler(&http.Client{}, lc, db)
	handler := http.HandlerFunc(h.BookingHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	db.AssertExpectations(t)
}
