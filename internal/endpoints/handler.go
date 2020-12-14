package endpoints

import (
	"net/http"

	"github.com/Sacro/SpaceTrouble/internal/repository"
	"github.com/Sacro/SpaceTrouble/internal/spacex"
)

type Handler struct {
	client     *http.Client
	repository *repository.Repository
}

func NewHandler(c *http.Client, r *repository.Repository) *Handler {
	return &Handler{
		client:     c,
		repository: r,
	}
}

func (h *Handler) getLaunches() (spacex.LaunchPads, error) {
	return spacex.GetLaunches(h.client)
}
