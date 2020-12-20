package spacex

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	c := HTTPLaunchClient{
		Client: client,
	}

	launches, err := c.GetLaunches(context.Background())
	assert.Nil(t, err)

	assert.NotNil(t, launches)
}
