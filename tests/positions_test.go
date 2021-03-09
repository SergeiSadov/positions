package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/sergeisadov/positions/internal/entities"
	"github.com/sergeisadov/positions/internal/service/http_api/responses"

	"github.com/valyala/fasthttp"
)

func testPositions(paramsStr string) (response responses.Positions, err error) {
	r, err := http.NewRequest(fasthttp.MethodGet, fmt.Sprintf("http://%s:%s/positions%s", cfg.Host, cfg.Port, paramsStr), nil)
	if err != nil {
		return response, err
	}

	res, err := serve(controller.Positions, r)
	if err != nil {
		return response, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	if err = json.Unmarshal(body, &response); err != nil {
		return response, err
	}

	return
}

func TestPositions(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    responses.Positions
		wantErr bool
	}{
		{"no domain", "", responses.Positions{}, true},
		{"empty page", "?domain=qwert.net&page=1", responses.Positions{Domain: "qwert.net", Positions: []entities.Position{}}, false},
		{"qwert.net", "?domain=qwert.net", responses.Positions{"qwert.net", []entities.Position{{
			Keyword:  "qwert",
			Position: 2,
			Domain:   "qwert.net",
			Url:      "https://qwert.net/bt.html",
			Volume:   37460000,
			Results:  270000000,
			Cpc:      25.76,
			Updated:  time.Date(2017, 05, 20, 02, 54, 07, 0, time.UTC),
		}}}, false},
		{"test.org", "?domain=test.org&page=0&sort=position", responses.Positions{"test.org", []entities.Position{{
			Keyword:  "t",
			Position: 4,
			Domain:   "test.org",
			Url:      "https://test.org/test.html",
			Volume:   37460000,
			Results:  270000000,
			Cpc:      25.76,
			Updated:  time.Date(2017, 05, 20, 02, 54, 07, 0, time.UTC),
		}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testPositions(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			if len(got.Positions) > 1 {
				got.Positions = []entities.Position{got.Positions[0]}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n got = %+v \n want %+v", got, tt.want)
			}

		})
	}
}
