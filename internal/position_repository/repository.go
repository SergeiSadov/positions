package position_repository

import "github.com/sergeisadov/positions/internal/entities"

type IPositionRepository interface {
	Count(Filters) (total int, err error)
	List(Filters) (positions []entities.Position, err error)
}

type Filters struct {
	Domain    string
	Page      int
	SortField string
	Order     string
}
