package postgresql

import (
	"database/sql"
	"filmlib/internal/config"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const nodeName = "storage"

func GetDBConnection(conf config.DatabaseConfig) (*sql.DB, error) {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?&search_path=%s&connect_timeout=%d",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
		// conf.AppName,
		conf.Schema,
		conf.ConnectionTimeout,
	)

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
