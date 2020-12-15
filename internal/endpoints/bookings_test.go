package endpoints

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sacro/SpaceTrouble/internal/repository"
	"github.com/Sacro/SpaceTrouble/internal/ticket"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestBookingsEndpointNoResults(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/bookings", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	db := &repository.MockRepository{}
	db.On("Bookings").Return([]ticket.Ticket{}, nil)

	h := NewHandler(&http.Client{}, db)
	handler := http.HandlerFunc(h.BookingsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response []ticket.Ticket
	err = json.NewDecoder(rr.Result().Body).Decode(&response)
	assert.Nil(t, err)

	assert.Equal(t, len(response), 0)

	db.AssertExpectations(t)
}

func TestBookingsEndpointOneResult(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/bookings", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	db := &repository.MockRepository{}
	f := faker.New()

	db.On("Bookings").Return([]ticket.Ticket{
		*ticket.Fake(f),
	}, nil)

	h := NewHandler(&http.Client{}, db)
	handler := http.HandlerFunc(h.BookingsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response []ticket.Ticket
	err = json.NewDecoder(rr.Result().Body).Decode(&response)
	assert.Nil(t, err)

	assert.Equal(t, len(response), 1)

	db.AssertExpectations(t)
}

func TestBookingsEndpointMultipleResults(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/bookings", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	db := &repository.MockRepository{}
	f := faker.New()

	db.On("Bookings").Return([]ticket.Ticket{
		*ticket.Fake(f),
		*ticket.Fake(f),
		*ticket.Fake(f),
	}, nil)

	h := NewHandler(&http.Client{}, db)
	handler := http.HandlerFunc(h.BookingsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response []ticket.Ticket
	err = json.NewDecoder(rr.Result().Body).Decode(&response)
	assert.Nil(t, err)

	assert.Equal(t, len(response), 3)

	db.AssertExpectations(t)
}
