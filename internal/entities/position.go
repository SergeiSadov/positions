package entities

import "time"

type Position struct {
	Keyword  string    `json:"keyword"`
	Position int64     `json:"position"`
	Domain   string    `json:"domain,omitempty"`
	Url      string    `json:"url"`
	Volume   int64     `json:"volume"`
	Results  int64     `json:"results"`
	Cpc      float64   `json:"cpc,omitempty"`
	Updated  time.Time `json:"updated"`
}
