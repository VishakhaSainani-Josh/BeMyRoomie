package users

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/models"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/errhandler"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/response"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/validations"
)

// Reads user details then extracts role - finder or lister from request url and Regsiters user using register service
func UserRegister(userService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}

		var user models.NewUserRequest
		err = json.Unmarshal(body, &user)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}

		err = validations.ValidateRegisterUserStruct(user)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}

		getRole := r.URL.Path
		role := strings.Split(getRole, "/")
		userId, err := userService.RegisterUser(ctx, user, role[1])
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}

		response.HandleResponse(w, http.StatusOK, "User Registered Successfully", userId)
	}

}

// Reads user email and password, and calls login service to allow a vaild user to log in
func LoginUser(userService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}

		var loginRequest models.LoginRequest
		err = json.Unmarshal(body, &loginRequest)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}
		err = validations.ValidateLoginRequestStruct(loginRequest)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}

		token, user, err := userService.LoginUser(ctx, loginRequest)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}

		resp := models.LoginResponse{Token: token, User: user}
		response.HandleResponse(w, http.StatusOK, "Signed in successfully", resp)
	}
}

// Reads user Preferences and makes request to add preference service
func AddPreferences(userService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}

		var preferences models.NewPreferenceRequest
		err = json.Unmarshal(body, &preferences)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}

		err = validations.ValidatePreferenceStruct(preferences)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}

		err = userService.AddPreferences(ctx, preferences)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}

		response.HandleResponse(w, http.StatusOK, "Preferences Updated Successfully", nil)
	}
}

func ViewProfile(userService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user, err := userService.ViewProfile(ctx)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}

		response.HandleResponse(w, http.StatusOK, "fetch profile details", user)
	}
}

func UpdateProfile(userService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}

		var user models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}

		err = validations.ValidateUpdateProfileStruct(user)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}

		err = userService.UpdateProfile(ctx, user)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage, nil)
			return
		}

		response.HandleResponse(w, http.StatusOK, "Profile Details Updated Successfully", nil)
	}
}
