package endpoints

import (
	"net/http"

	"github.com/Sacro/SpaceTrouble/internal/repository"
	"github.com/Sacro/SpaceTrouble/internal/spacex"
)

type LaunchHandler interface {
	getLaunches() (spacex.LaunchPads, error)
}

type Handler struct {
	client       *http.Client
	launchClient spacex.LaunchClient
	repository   repository.TicketRepository
}

func NewHandler(c *http.Client, lc spacex.LaunchClient, r repository.TicketRepository) *Handler {
	return &Handler{
		client:       c,
		launchClient: lc,
		repository:   r,
	}
}
