package app

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/app/users"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/models"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/errhandler"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/response"
)

const (
	bodyError  = "error reading body"
	invalidReq = "error invalid request"
)

func UserRegister(userService users.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.HandleResponse(w, http.StatusInternalServerError, bodyError)
			return
		}

		var user models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			response.HandleResponse(w, http.StatusBadRequest, invalidReq)
			return
		}

		getRole := r.URL.Path
		role := strings.Split(getRole, "/")
		userId, err := userService.RegisterUser(user, role[1])
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		resp := struct {
			UserId int    `json:"user_id"`
			Msg    string `json:"msg"`
		}{UserId: userId, Msg: "User Registered Successfully"}

		response.HandleResponse(w, http.StatusOK, resp)
	}

}

func LoginUser(userService users.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.HandleResponse(w, http.StatusInternalServerError, bodyError)
			return
		}

		var loginRequest struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		err = json.Unmarshal(body, &loginRequest)
		if err != nil {
			response.HandleResponse(w, http.StatusBadRequest, invalidReq)
			return
		}

		token, err := userService.LoginUser(loginRequest.Email, loginRequest.Password)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		resp := struct {
			Token string `json:"token"`
			Msg   string `json:"msg"`
		}{Token: token, Msg: "Signed in successfully"}

		response.HandleResponse(w, http.StatusOK, resp)
	}
}
