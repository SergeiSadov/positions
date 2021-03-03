package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/valyala/fasthttp"

	"github.com/sergeisadov/positions/internal/config"
	di "github.com/sergeisadov/positions/internal/di"
	"github.com/sergeisadov/positions/internal/service/http_api"
)

func main() {
	var pathConfig string
	flag.StringVar(&pathConfig, "config", "config.json", "")
	flag.Parse()

	cfg, err := config.Load(pathConfig)
	if err != nil {
		log.Fatal(err)
	}

	container, err := di.Register(cfg)
	if err != nil {
		log.Fatal(err)
	}

	server := &fasthttp.Server{
		Handler: http_api.New(container).Handler,
	}

	go func() {
		if err = server.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)); err != nil {
			container.Logger.Fatal().Err(err)
		}
	}()

	container.Logger.Info().Msg("app started")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan

	container.Logger.Info().Str("signal", sig.String()).Msg("starting graceful shutdown...")

	if err = server.Shutdown(); err != nil {
		container.Logger.Info().Str("signal", sig.String()).Err(err).Send()
	}

	container.Logger.Info().Msg("app stopped")
}
