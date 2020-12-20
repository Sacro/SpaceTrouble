package spacex

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type LaunchClient interface {
	GetLaunches(context.Context) (LaunchPads, error)
}

type HttpLaunchClient struct {
	http.Client
}

type MockLaunchClient struct {
	mock.Mock
}

// Ensure the implementation matches the interface
var _ LaunchClient = &HttpLaunchClient{}

func (c *HttpLaunchClient) GetLaunches(ctx context.Context) (LaunchPads, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.spacexdata.com/v4/launches", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var launches LaunchPads
	err = json.NewDecoder(res.Body).Decode(&launches)

	if err != nil {
		return nil, err
	}

	return launches, nil
}

// Ensure the implementation matches the interface
var _ LaunchClient = &MockLaunchClient{}

func (m *MockLaunchClient) GetLaunches(ctx context.Context) (LaunchPads, error) {
	args := m.Called()

	return args.Get(0).(LaunchPads), args.Error(1)
}
