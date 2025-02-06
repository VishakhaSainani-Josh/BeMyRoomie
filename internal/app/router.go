package app

import (
	"net/http"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/app/users"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/middleware"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/response"
)

func InitRouter(deps dependencies) *http.ServeMux {
	router := http.NewServeMux()

	//Public routes
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		response.HandleResponse(w, http.StatusOK, "Server working")
	})

	router.HandleFunc("POST /finder/signup", users.UserRegister(deps.userService))
	router.HandleFunc("POST /lister/signup", users.UserRegister(deps.userService))

	router.HandleFunc("POST /signin", users.LoginUser(deps.userService))

	//Protected routes
	router.HandleFunc("POST /preferences", middleware.AuthMiddleware(users.AddPreferences(deps.userService)))

	return router
}
