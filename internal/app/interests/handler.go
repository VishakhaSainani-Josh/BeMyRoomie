package interests

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/models"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/errhandler"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/response"
)

// Retreives property id from the url and allows a user to add interest by calling show interest service
func ShowInterest(interestService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		getPropertyId := r.URL.Path
		propertyIdPath := strings.Split(getPropertyId, "/")
		propertyId, err := strconv.Atoi(propertyIdPath[2])
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		err = interestService.ShowInterest(ctx, propertyId)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		response.HandleResponse(w, http.StatusOK, "Interest shown successfully")
	}
}

// Allows a user to check properties they added interest to, by calling get interested properties service
func GetInterestedProperties(interestService Service) func(e http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		properties, err := interestService.GetInterestedProperties(ctx)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}
		
		response.HandleResponse(w, http.StatusOK, properties)
	}
}

// Retreives property id from url and allows a user to remove interest from a property by calling remove interest service
func RemoveInterest(interestService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		getPropertyId := r.URL.Path
		propertyIdPath := strings.Split(getPropertyId, "/")
		propertyId, err := strconv.Atoi(propertyIdPath[3])
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		err = interestService.RemoveInterest(ctx, propertyId)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		response.HandleResponse(w, http.StatusOK, "Interest deleted successfully")
	}
}

/*
Retreives property id from url and allows a lister to get users who shown interest in their listed property by
calling get interested users service
*/
func GetInterestedUsers(interestService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		getPropertyId := r.URL.Path
		propertyIdPath := strings.Split(getPropertyId, "/")
		propertyId, err := strconv.Atoi(propertyIdPath[3])
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		user, err := interestService.GetInterestedUsers(ctx, propertyId)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		response.HandleResponse(w, http.StatusOK, user)
	}
}

// Reads interest status from body whether accept or rejcet and updates it by calling accept interest service
func AcceptInterest(interestService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		var interestStatus models.InterestStatusRequest
		err = json.Unmarshal(body, &interestStatus)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		requestURL := r.URL.Path
		urlArr := strings.Split(requestURL, "/")
		propertyId, err := strconv.Atoi(urlArr[3])
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		userId, err := strconv.Atoi(urlArr[4])
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		err = interestService.AcceptInterest(ctx, userId, propertyId, interestStatus.IsAccepted)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		response.HandleResponse(w, http.StatusOK, "Status Updated successfully")
	}
}
