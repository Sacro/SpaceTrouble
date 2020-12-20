package endpoints

import (
	"github.com/Sacro/SpaceTrouble/internal/spacex"
	"github.com/stretchr/testify/mock"
)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) getLaunches() (spacex.LaunchPads, error) {
	args := m.Called()

	return args.Get(0).(spacex.LaunchPads), args.Error(1)
}

var _ LaunchHandler = &MockHandler{}
