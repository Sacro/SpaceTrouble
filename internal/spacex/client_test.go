package spacex

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	launches, err := GetLaunches(context.Background(), client)
	assert.Nil(t, err)

	assert.NotNil(t, launches)
}
