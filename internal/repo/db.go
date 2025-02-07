package repo

import (
	"database/sql"
	"errors"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/config"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// Databse Connection
func ConnectDB() (*sql.DB, error) {
	config.Load()

	dbURL := viper.GetString("POSTGRESQL_URL")
	if dbURL == "" {
		return nil, errors.New("POSTGRESQL_URL not found")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
