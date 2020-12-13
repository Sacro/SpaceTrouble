package ticket

import (
	"time"

	"github.com/jaswdr/faker"
)

// Ticket contains the data for booking a ticket
type Ticket struct {
	FirstName     string    `json:"first_name" validate:"required"`
	LastName      string    `json:"last_name" validate:"required"`
	Gender        string    `json:"gender" validate:"required"`
	Birthday      time.Time `json:"birthday" validate:"required,lt"`
	LaunchpadID   string    `json:"launchpad_id" validate:"required"`
	DestinationID string    `json:"destination_id" `
	LaunchDate    time.Time `json:"launch_date" validate:"required,gt"`
}

func Fake(f faker.Faker) Ticket {
	return Ticket{
		FirstName:     f.Person().FirstName(),
		LastName:      f.Person().LastName(),
		Gender:        f.Person().Title(),
		Birthday:      f.Time().Time(time.Now().UTC()),
		LaunchpadID:   f.UUID().V4(),
		DestinationID: f.UUID().V4(),
		LaunchDate:    time.Now().Add(time.Hour * 24),
	}
}
