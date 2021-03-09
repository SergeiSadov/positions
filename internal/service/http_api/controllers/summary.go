package controllers

import (
	"encoding/json"

	"github.com/sergeisadov/positions/internal/position_repository"
	"github.com/sergeisadov/positions/internal/service/http_api/responses"

	"github.com/valyala/fasthttp"
)

func (c *Controller) Summary(ctx *fasthttp.RequestCtx) {
	domain := string(ctx.QueryArgs().Peek(domainParam))
	if domain == "" {
		c.BadRequest(ctx)
		return
	}

	total, err := c.repo.Count(position_repository.Filters{Domain: domain})
	if err != nil {
		c.logger.Info().Msg("Hello World!")

		c.logger.Error().Str("module", "positions controller").Str("action", "get count").Err(err).Msg("")
		c.BadRequest(ctx)
		return
	}

	summary := responses.Summary{
		Domain:         domain,
		PositionsCount: total,
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
