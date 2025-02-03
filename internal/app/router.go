package app

import (
	"net/http"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/response"
)

func InitRouter(deps dependencies) *http.ServeMux {
	router := http.NewServeMux()

	//Open routes
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		response.HandleResponse(w, http.StatusOK, "Server working")
	})

	router.HandleFunc("POST /finder/signup", UserRegister(deps.userService))
	router.HandleFunc("POST /lister/signup", UserRegister(deps.userService))

	router.HandleFunc("POST /signin", LoginUser(deps.userService))

	return router
}
