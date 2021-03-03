package controllers

import (
	"github.com/rs/zerolog"

	repository "github.com/sergeisadov/positions/internal/position_repository"
)

const (
	domainParam    = "domain"
	pageParam      = "page"
	sortFieldParam = "sortField"
)

type Controller struct {
	repo   repository.IPositionRepository
	logger *zerolog.Logger
}

func New(repo repository.IPositionRepository, logger *zerolog.Logger) *Controller {
	return &Controller{repo: repo, logger: logger}
}
