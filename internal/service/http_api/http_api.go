package http_api

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	"github.com/sergeisadov/positions/internal/di"
	"github.com/sergeisadov/positions/internal/service/http_api/controllers"

	"github.com/fasthttp/router"
)

func New(di *di.Di) (mux *router.Router) {
	mux = router.New()

	controller := controllers.New(di.Repository, &di.Logger)

	mux.PanicHandler = controller.PanicHandler
	mux.NotFound = controller.NotFound
	mux.MethodNotAllowed = controller.MethodNotAllowed

	mux.GET("/positions", controller.Positions)
	mux.GET("/summary", controller.Summary)

	mux.GET("/info", func(ctx *fasthttp.RequestCtx) {
	})

	mux.GET("/metrics", fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler()))

	return mux
}
