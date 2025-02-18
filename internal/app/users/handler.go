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
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		var user models.NewUserRequest
		err = json.Unmarshal(body, &user)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		err = validations.ValidateRegisterUserStruct(user)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		getRole := r.URL.Path
		role := strings.Split(getRole, "/")
		userId, err := userService.RegisterUser(ctx, user, role[1])
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		resp := models.UserResponse{UserId: userId, Message: "User Registered Successfully"}
		response.HandleResponse(w, http.StatusOK, resp)
	}

}

// Reads user email and password, and calls login service to allow a vaild user to log in
func LoginUser(userService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		var loginRequest models.LoginRequest
		err = json.Unmarshal(body, &loginRequest)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}
		err = validations.ValidateLoginRequestStruct(loginRequest)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		token, err := userService.LoginUser(ctx, loginRequest)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		resp := models.LoginResponse{Token: token, Message: "Signed in successfully"}
		response.HandleResponse(w, http.StatusOK, resp)
	}
}

// Reads user Preferences and makes request to add preference service
func AddPreferences(userService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		var preferences models.NewPreferenceRequest
		err = json.Unmarshal(body, &preferences)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		err = validations.ValidatePreferenceStruct(preferences)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		err = userService.AddPreferences(ctx, preferences)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		response.HandleResponse(w, http.StatusOK, "Preferences Updated Successfully")
	}
}

func ViewProfile(userService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user, err := userService.ViewProfile(ctx)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		response.HandleResponse(w, http.StatusOK, user)
	}
}

func UpdateProfile(userService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		var user models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		err = validations.ValidateUpdateProfileStruct(user)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		err = userService.UpdateProfile(ctx, user)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		response.HandleResponse(w, http.StatusOK, "Profile Details Updated Successfully")
	}
}
