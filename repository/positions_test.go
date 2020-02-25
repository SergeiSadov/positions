package repository

import (
	"github.com/SergeiSadov/positions/config"
	"github.com/SergeiSadov/positions/db"
	"log"
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

func TestCountPositionsByDomain(t *testing.T) {
	precondition()

	type args struct {
		domain string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int64
		wantErr   bool
	}{
		{"1", args{domain: "abesse.net"}, 238, false},
		{"2", args{domain: "test"}, 0, false},
		{"3", args{domain: ""}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, err := CountPositionsByDomain(tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountPositionsByDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCount != tt.wantCount {
				t.Errorf("CountPositionsByDomain() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestGetAllByDomain(t *testing.T) {
	precondition()

	type args struct {
		domain    string
		sortField string
		page      int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"1", args{
			domain:    "abesse.net",
			sortField: "Value",
			page:      1,
		}, false},
		{"2", args{
			domain:    "",
			sortField: "",
			page:      0,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetAllByDomain(tt.args.domain, tt.args.sortField, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllByDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
