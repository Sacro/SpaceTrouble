package repository

import (
	"github.com/Sacro/SpaceTrouble/internal/ticket"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateBooking(ticket *ticket.Ticket) error {
	args := m.Called(ticket)

	return args.Error(0)
}

func (m *MockRepository) Bookings() ([]ticket.Ticket, error) {
	args := m.Called()

	return args.Get(0).([]ticket.Ticket), args.Error(1)
}

var _ TicketRepository = &MockRepository{}
