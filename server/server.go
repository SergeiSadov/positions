package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/SergeiSadov/positions/api"
	"github.com/SergeiSadov/positions/config"
	"github.com/SergeiSadov/positions/db"
	"github.com/SergeiSadov/positions/elasticlog"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start() error {
	var pathConfig string
	flag.StringVar(&pathConfig, "config", "config.json", "")
	flag.Parse()

	err := config.Load(pathConfig)
	if err != nil {
		return fmt.Errorf("server.Start() error #1: \t\n %w \r\n", err)
	}
	fmt.Println("loaded config")

	err = elasticlog.StartLog()
	if err != nil {
		return fmt.Errorf("server.Start() error #2: \t\n %w \r\n", err)
	}
	fmt.Println("log started")

	err = db.Connect(config.Config.DbDriver, config.Config.DbPath)
	if err != nil {
		return fmt.Errorf("server.Start() error #3: \t\n %w \r\n", err)
	}

	defer func() {
		err = db.DB.Close()
		if err != nil {
			log.Printf("server.Start() error #4: \t\n %v \r\n", err.Error())
		}
	}()

	fmt.Println("database connected")

	r := mux.NewRouter()

	summaryRequest := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "summary",
		})

	prometheus.MustRegister(summaryRequest)

	positionsRequest := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "positions",
		})

	prometheus.MustRegister(positionsRequest)

	r.HandleFunc("/summary", func(w http.ResponseWriter, r *http.Request) {
		summaryRequest.Inc()
		api.Summary(w, r)
	}).Methods(http.MethodGet)
	r.HandleFunc("/positions", func(w http.ResponseWriter, r *http.Request) {
		positionsRequest.Inc()
		api.Position(w, r)
	}).Methods(http.MethodGet)

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
	r.Handle("/metrics", promhttp.Handler())

	fmt.Printf("listening port %s \r\n", config.Config.Port)

	return http.ListenAndServe(fmt.Sprintf(":%s", config.Config.Port), r)
}
