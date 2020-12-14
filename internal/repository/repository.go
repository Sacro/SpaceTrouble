package repository

import (
	"github.com/Sacro/SpaceTrouble/internal/ticket"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Repository struct {
	db *pg.DB
}

func New(db *pg.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) CreateSchema() error {
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

func (repo *Repository) CreateBooking(ticket *ticket.Ticket) error {
	_, err := repo.db.Model(ticket).Insert()

	return err
}

func (repo *Repository) Bookings() ([]ticket.Ticket, error) {
	var tickets []ticket.Ticket

	if err := repo.db.Model(&tickets).Select(); err != nil {
		return nil, err
	}

	if tickets == nil {
		tickets = make([]ticket.Ticket, 0)
	}

	return tickets, nil
}
