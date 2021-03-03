package responses

import "github.com/sergeisadov/positions/internal/entities"

type Positions struct {
	Domain    string              `json:"domain"`
	Positions []entities.Position `json:"positions"`
}
