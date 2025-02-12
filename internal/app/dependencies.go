package app

import (
	"database/sql"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/app/properties"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/app/users"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/repo"
)

type dependencies struct {
	userService     users.Service
	propertyService properties.Service
}

// Start services
func InitServices(db *sql.DB) dependencies {
	userRepo := repo.NewUserRepo(db)
	userService := users.NewService(userRepo)

	propertyRepo := repo.NewPropertyRepo(db)
	propertyService := properties.NewService(propertyRepo)

	return dependencies{userService: userService,
		propertyService: propertyService}

}
