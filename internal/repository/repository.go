package repository

import (
	"github.com/Sacro/SpaceTrouble/internal/ticket"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type TicketRepository interface {
	CreateBooking(*ticket.Ticket) error
	Bookings() ([]ticket.Ticket, error)
}

// Ensure the implementation matches the interface
var _ TicketRepository = &TicketRepo{}

type TicketRepo struct {
	db *pg.DB
}

func New(db *pg.DB) *TicketRepo {
	return &TicketRepo{
		db: db,
	}
}

func (repo *TicketRepo) CreateSchema() error {
	models := []interface{}{
		(*ticket.Ticket)(nil),
	}

	for _, model := range models {
		err := repo.db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *TicketRepo) CreateBooking(t *ticket.Ticket) error {
	_, err := repo.db.Model(t).Insert()

	return err
}

func (repo *TicketRepo) Bookings() ([]ticket.Ticket, error) {
	var tickets []ticket.Ticket

	if err := repo.db.Model(&tickets).Select(); err != nil {
		return nil, err
	}

	if tickets == nil {
		tickets = make([]ticket.Ticket, 0)
	}

	return tickets, nil
}
