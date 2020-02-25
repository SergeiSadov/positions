package api

import (
	"encoding/json"
	"net/http"

	"github.com/SergeiSadov/positions/elasticlog"
	"github.com/SergeiSadov/positions/repository"
)

type summaryResponse struct {
	Domain         string `json:"domainParam"`
	PositionsCount int64  `json:"positions_count"`
}

func Summary(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get(domainParam)
	if domain == "" {
		elasticlog.Log.Errorf("api.Summary() error #1: \t\n %v \r\n", "empty domainParam")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	count, err := repository.CountPositionsByDomain(domain)
	if err != nil {
		elasticlog.Log.Errorf("api.Summary() error #2: \t\n %v \r\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := json.Marshal(summaryResponse{
		Domain:         domain,
		PositionsCount: count,
	})
	if err != nil {
		elasticlog.Log.Errorf("api.Summary() error #3: \t\n %v \r\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(result)
	if err != nil {
		elasticlog.Log.Errorf("api.Summary() error #4: \t\n %v \r\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
