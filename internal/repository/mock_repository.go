package repository

import (
	"github.com/Sacro/SpaceTrouble/internal/ticket"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateBooking(*ticket.Ticket) error {
	return nil
}

func (m *MockRepository) Bookings() ([]ticket.Ticket, error) {
	return nil, nil
}

var _ TicketRepository = &MockRepository{}
