package di

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/rs/zerolog"

	_ "github.com/mattn/go-sqlite3"

	"github.com/sergeisadov/positions/internal/config"
	positionRepository "github.com/sergeisadov/positions/internal/position_repository"
	"github.com/sergeisadov/positions/internal/position_repository/sqlite"
)

type Di struct {
	Logger     zerolog.Logger
	Repository positionRepository.IPositionRepository
	Config     config.Config
	DB         *sql.DB
}

func Register(config config.Config) (*Di, error) {
	db, err := sql.Open(config.Database.Driver, config.Database.Path)
	if err != nil {
		return nil, fmt.Errorf("error opening db connection: %w", err)
	}

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	multi := zerolog.MultiLevelWriter(consoleWriter, os.Stdout)

	return &Di{
		Logger:     zerolog.New(multi).With().Timestamp().Logger(),
		Repository: sqlite.New(db),
		DB:         db,
	}, nil
}
