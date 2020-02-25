package api

import (
	"github.com/SergeiSadov/positions/config"
	"github.com/SergeiSadov/positions/db"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func precondition() {
	err := config.Load("../config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Connect(config.Config.DbDriver, config.Config.DbPath)
	if err != nil {
		log.Fatal(err)
	}
}

func TestPosition(t *testing.T) {
	precondition()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{"1", args{
			w: httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet, "/positions?domain=abesse.net&page=1", nil),
		},
			200,
		},
		{"2", args{
			w: httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet, "/positions?domain=", nil),
		},
			400,
		},
		{"3", args{
			w: httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet, "/positions?domain=test", nil),
		},
			200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Position(tt.args.w, tt.args.r)
			if tt.args.w.(*httptest.ResponseRecorder).Code != tt.wantCode {
				t.Errorf("wrong code! expected: %d, actual: %d\n\n", tt.wantCode,
					tt.args.w.(*httptest.ResponseRecorder).Code)
			}
		})
	}
}

func TestSummary(t *testing.T) {
	precondition()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{"1", args{
			w: httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet, "/summary?domain=abesse.net", nil),
		},
			200,
		},
		{"2", args{
			w: httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet, "/summary?domain=", nil),
		},
			400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Summary(tt.args.w, tt.args.r)
			if tt.args.w.(*httptest.ResponseRecorder).Code != tt.wantCode {
				t.Errorf("wrong code! expected: %d, actual: %d\n\n", tt.wantCode,
					tt.args.w.(*httptest.ResponseRecorder).Code)
			}
		})
	}
}
