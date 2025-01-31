package repo

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/config"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *sql.DB

func ConnectDB() {
	config.Load()

	dbURL := viper.GetString("POSTGRESQL_URL")
	if dbURL == "" {
		log.Fatal("POSTGRESQL_URL not found")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database connection failed: %s", err)
	}

	fmt.Println("Connected to Database")
	DB = db
}
