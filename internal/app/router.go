package app

import (
	"net/http"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/app/properties"
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
	router.HandleFunc("POST /preferences", middleware.Authentication(users.AddPreferences(deps.userService)))
	router.HandleFunc("POST /property", middleware.Authentication(middleware.AuthorizeLister(properties.RegisterProperty(deps.propertyService))))
	router.HandleFunc("PATCH /properties/{property_id}", middleware.Authentication(middleware.AuthorizeLister(properties.UpdateProperty(deps.propertyService))))
	router.HandleFunc("GET /properties", middleware.Authentication(properties.GetAllProperties(deps.propertyService)))
	router.HandleFunc("GET /properties/{property_id}", middleware.Authentication(properties.GetParticularProperties(deps.propertyService)))

	return router
}
