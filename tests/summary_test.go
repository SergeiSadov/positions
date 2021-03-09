package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/sergeisadov/positions/internal/service/http_api/responses"

	"github.com/valyala/fasthttp"
)

func testSummary(paramsStr string) (response responses.Summary, err error) {
	r, err := http.NewRequest(fasthttp.MethodGet, fmt.Sprintf("http://%s:%s/summary%s", cfg.Host, cfg.Port, paramsStr), nil)
	if err != nil {
		return response, err
	}

	res, err := serve(controller.Summary, r)
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

func TestSummary(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    responses.Summary
		wantErr bool
	}{
		{"no domain", "", responses.Summary{}, true},
		{"test.org", "?domain=test.org", responses.Summary{"test.org", 3}, false},
		{"qwert.net", "?domain=qwert.net", responses.Summary{"qwert.net", 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testSummary(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}

		})
	}
}
