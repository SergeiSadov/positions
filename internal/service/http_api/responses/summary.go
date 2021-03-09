package responses

type Summary struct {
	Domain         string `json:"domain"`
	PositionsCount int    `json:"positions_count"`
}
