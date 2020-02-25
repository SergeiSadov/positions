package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SergeiSadov/positions/entities"
	"github.com/SergeiSadov/positions/repository"

	"github.com/SergeiSadov/positions/elasticlog"
)

type positionResponse struct {
	Domain    string              `json:"domain"`
	Positions []entities.Position `json:"positions"`
}

func Position(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	domain := query.Get(domainParam)
	if domain == "" {
		elasticlog.Log.Errorf("api.Position() error #1: \t\n %v \r\n", "empty domainParam")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sortField := query.Get(sortFieldParam)
	if sortField == "" {
		sortField = defaultSortField
	}

	page := 0
	pageParam := query.Get(pageParam)
	if pageParam != "" {
		pageNumber, err := strconv.Atoi(pageParam)
		if err != nil {
			elasticlog.Log.Errorf("api.Position() error #2: \t\n %v \r\n", err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
		page = pageNumber
	}

	positions, err := repository.GetAllByDomain(domain, sortField, page)
	if err != nil {
		elasticlog.Log.Errorf("api.Position() error #3: \t\n %v \r\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := json.Marshal(positionResponse{
		Domain:    domain,
		Positions: positions,
	})
	if err != nil {
		elasticlog.Log.Errorf("api.Position() error #4: \t\n %v \r\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(result)
	if err != nil {
		elasticlog.Log.Errorf("api.Position() error #5: \t\n %v \r\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
