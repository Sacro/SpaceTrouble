package spacex

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	launches, err := GetLaunches(client)
	assert.Nil(t, err)

	assert.NotNil(t, launches)
}
