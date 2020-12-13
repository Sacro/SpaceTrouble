package ticket

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestFakeTicketIsValid(t *testing.T) {
	f := faker.New()
	v := validator.New()

	ticket := Fake(f)
	err := v.Struct(ticket)
	assert.NoError(t, err)
}
