package repository

import (
	"fmt"
	"github.com/SergeiSadov/positions/db"
	"github.com/SergeiSadov/positions/entities"
)

const pageSize = 10

func CountPositionsByDomain(domain string) (count int64, err error) {
	return count, db.DB.QueryRow(`
SELECT COUNT()
FROM positions
WHERE domain = ?
`, domain).Scan(&count)
}

func GetAllByDomain(domain string, sortField string, page int) (positions []entities.Position, err error) {
	rows, err := db.DB.Query(`
SELECT keyword,
       position,
       url,
       volume,
       results,
       cpc,
       updated
FROM positions
WHERE domain = ?
ORDER BY ? LIMIT ?, ?
`, domain, sortField, page*pageSize, pageSize)
	if err != nil {
		return positions, fmt.Errorf("repository.GetAllByDomain() error #2: \t\n %w \r\n", err)
	}

	for rows.Next() {
		var position entities.Position
		err = rows.Scan(
			&position.Keyword,
			&position.Position,
			&position.Url,
			&position.Volume,
			&position.Results,
			&position.Cpc,
			&position.Updated,
		)
		if err != nil {
			return positions, fmt.Errorf("repository.GetAllByDomain() error #2: \t\n %w \r\n", err)
		}

		positions = append(positions, position)
	}

	return
}
