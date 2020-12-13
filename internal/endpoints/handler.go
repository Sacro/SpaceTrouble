package endpoints

import "github.com/Sacro/SpaceTrouble/internal/repository"

type Handler struct {
	repository *repository.Repository
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{
		repository: r,
	}
}
