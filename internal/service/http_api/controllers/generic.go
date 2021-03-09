package controllers

import (
	"github.com/valyala/fasthttp"
)

func (c *Controller) PanicHandler(ctx *fasthttp.RequestCtx, i interface{}) {
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	ctx.SetBodyString(fasthttp.StatusMessage(fasthttp.StatusInternalServerError))
}

func (c *Controller) NotFound(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusNotFound)
	ctx.SetBodyString(fasthttp.StatusMessage(fasthttp.StatusNotFound))
}

func (c *Controller) BadRequest(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusBadRequest)
	ctx.SetBodyString(fasthttp.StatusMessage(fasthttp.StatusBadRequest))
}

func (c *Controller) InternalError(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	ctx.SetBodyString(fasthttp.StatusMessage(fasthttp.StatusInternalServerError))
}

func (c *Controller) MethodNotAllowed(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
	ctx.SetBodyString(fasthttp.StatusMessage(fasthttp.StatusMethodNotAllowed))
}
