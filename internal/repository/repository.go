package repository

import (
	"github.com/Sacro/SpaceTrouble/internal/ticket"
	"github.com/go-pg/pg/v10"
)

type Repository struct {
	db *pg.DB
}

func New(db *pg.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) Booking(id string) (*ticket.Ticket, error) {
	ticket := &ticket.Ticket{
		ID: id,
	}

	if err := repo.db.Model(ticket).WherePK().Select(); err != nil {
		return ticket, err
	}

	return ticket, nil
}

func (repo *Repository) Bookings() ([]ticket.Ticket, error) {
	var tickets []ticket.Ticket

	if err := repo.db.Model(&tickets).Select(); err != nil {
		return nil, err
	}

	return tickets, nil
}
