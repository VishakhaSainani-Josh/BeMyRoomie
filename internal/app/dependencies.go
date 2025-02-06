package app

import (
	"database/sql"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/app/users"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/repo"
)

type dependencies struct {
	userService users.Service
}

//Start services
func InitServices(db *sql.DB) dependencies {
	userRepo := repo.NewUserRepo(db)

	userService := users.NewService(userRepo)

	return dependencies{userService: userService}
}
