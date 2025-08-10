package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const driverName = "postgres"

type Settings struct {
	Host          string
	Port          string
	User          string
	Password      string
	DbName        string
	Schema        string
	AppName       string
	MaxIdleConns  int
	MaxOpenConns  int
}

func Connect(settings Settings) (*sqlx.DB, error) {
	addr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s application_name=%s timezone=UTC",
		settings.Host,
		settings.Port,
		settings.User,
		settings.Password,
		settings.DbName,
		settings.Schema,
		settings.AppName,
	)

	sqlxDB, err := sqlx.Connect(driverName, addr)
	if err != nil {
		return nil, err
	}
	sqlxDB.SetMaxIdleConns(settings.MaxIdleConns)
	sqlxDB.SetMaxOpenConns(settings.MaxOpenConns)

	return sqlxDB, nil
}
