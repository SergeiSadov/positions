package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/valyala/fasthttp"

	"github.com/sergeisadov/positions/internal/entities"
	"github.com/sergeisadov/positions/internal/position_repository"
	"github.com/sergeisadov/positions/internal/service/http_api/responses"
)

const (
	defaultSortField = "volume"
)

func (c *Controller) Positions(ctx *fasthttp.RequestCtx) {
	var filters position_repository.Filters
	domain := string(ctx.QueryArgs().Peek(domainParam))
	if domain == "" {
		c.BadRequest(ctx)
		return
	}

	filters.Domain = domain

	sort := string(ctx.QueryArgs().Peek(sortFieldParam))
	if sort == "" {
		filters.SortField = defaultSortField
	} else {
		filters.SortField = sort
	}

	pageStr := string(ctx.QueryArgs().Peek(pageParam))
	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			c.BadRequest(ctx)
			return
		}

		filters.Page = page
	}

	positions, err := c.repo.List(filters)
	if err != nil {
		c.logger.Error().Str("module", "positions controller").Str("action", "get count").Err(err).Msg("")
		c.BadRequest(ctx)
		return
	}

	if positions == nil {
		positions = []entities.Position{}
	}

	summary := responses.Positions{
		Domain:    domain,
		Positions: positions,
	}

	bytes, err := json.Marshal(summary)
	if err != nil {
		c.logger.Error().Str("module", "positions controller").Str("action", "get marshall json").Err(err)
		c.BadRequest(ctx)
		return
	}

	_, err = ctx.Write(bytes)
	if err != nil {
		c.logger.Error().Str("module", "positions controller").Str("action", "write response").Err(err)
		c.BadRequest(ctx)
		return
	}
}
