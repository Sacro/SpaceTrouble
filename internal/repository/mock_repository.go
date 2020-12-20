package repository

import (
	"github.com/Sacro/SpaceTrouble/internal/ticket"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateBooking(t *ticket.Ticket) error {
	args := m.Called(t)

	return args.Error(0)
}

func (m *MockRepository) Bookings() ([]ticket.Ticket, error) {
	args := m.Called()

	return args.Get(0).([]ticket.Ticket), args.Error(1)
}

func (m *MockRepository) DeleteBooking(id string) error {
	args := m.Called(id)

	return args.Error(0)
}

var _ TicketRepository = &MockRepository{}
