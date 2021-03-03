package sqlite

import (
	"database/sql"
	"fmt"

	repository "github.com/sergeisadov/positions/internal/position_repository"

	"github.com/sergeisadov/positions/internal/entities"
)

const (
	PositionsTable = "positions"
)

const limit = 30

type PositionsRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *PositionsRepository {
	return &PositionsRepository{db: db}
}

func (pr *PositionsRepository) Count(filters repository.Filters) (total int, err error) {
	query := `SELECT COUNT(*) FROM ` + PositionsTable
	query, args := pr.applyFilters(query, filters)
	if err = pr.db.QueryRow(query, args...).Scan(&total); err != nil {
		return total, err
	}

	return
}

func (pr *PositionsRepository) List(filters repository.Filters) (positions []entities.Position, err error) {
	query := `SELECT keyword,
	       position,
           domain,
	       url,
	       volume,
	       results,
	       cpc,
	       updated
	FROM ` + PositionsTable
	query, args := pr.applyFilters(query, filters)

	query += " ORDER BY ?"
	args = append(args, filters.SortField)

	query += " LIMIT ? OFFSET ? "
	if filters.Page != 0 {
		args = append(args, limit, filters.Page*limit)
	} else {
		args = append(args, limit, filters.Page)
	}

	rows, err := pr.db.Query(query, args...)
	if err != nil {
		return positions, fmt.Errorf("repository.GetAllByDomain() error #1: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var position entities.Position
		err = rows.Scan(
			&position.Keyword,
			&position.Position,
			&position.Domain,
			&position.Url,
			&position.Volume,
			&position.Results,
			&position.Cpc,
			&position.Updated,
		)
		if err != nil {
			return positions, fmt.Errorf("repository.GetAllByDomain() error #2: %w", err)
		}

		positions = append(positions, position)
	}

	return
}

func (pr *PositionsRepository) applyFilters(query string, filters repository.Filters) (string, []interface{}) {
	var args []interface{}

	query += " WHERE 1"
	if filters.Domain != "" {
		query += " AND domain = ?"
		args = append(args, filters.Domain)
	}

	return query, args
}
