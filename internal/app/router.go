package app

import (
	"net/http"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/app/interests"
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
	router.HandleFunc("GET /profile", middleware.Authentication(users.ViewProfile(deps.userService)))
	router.HandleFunc("PATCH /profile", middleware.Authentication(users.UpdateProfile(deps.userService)))
	router.HandleFunc("POST /property", middleware.Authentication(middleware.AuthorizeLister(properties.RegisterProperty(deps.propertyService))))
	router.HandleFunc("PATCH /properties/{property_id}", middleware.Authentication(middleware.AuthorizeLister(properties.UpdateProperty(deps.propertyService))))
	router.HandleFunc("GET /properties", middleware.Authentication(properties.GetAllProperties(deps.propertyService)))
	router.HandleFunc("GET /properties/{property_id}", middleware.Authentication(properties.GetParticularProperties(deps.propertyService)))

	router.HandleFunc("POST /interest/{property_id}", middleware.Authentication(middleware.AuthorizeFinder(interests.ShowInterest(deps.interestService))))
	router.HandleFunc("GET /interests/properties", middleware.Authentication(middleware.AuthorizeFinder(interests.GetInterestedProperties(deps.interestService))))
	router.HandleFunc("DELETE /interests/properties/{property_id}", middleware.Authentication(middleware.AuthorizeFinder(interests.RemoveInterest(deps.interestService))))

	router.HandleFunc("GET /interests/properties/{property_id}", middleware.Authentication(middleware.AuthorizeLister(interests.GetInterestedUsers(deps.interestService))))
	router.HandleFunc("PATCH /interests/properties/{property_id}/{user_id}", middleware.Authentication(middleware.AuthorizeLister(interests.AcceptInterest(deps.interestService))))
	return router
}
