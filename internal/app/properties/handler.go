package properties

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/models"
	constant "github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/constants"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/errhandler"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/response"
)

// Reads property details and calls register property service to allow a lister to post a property
func RegisterProperty(propertyService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		var property models.NewPropertyRequest
		err = json.Unmarshal(body, &property)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		propertyId, err := propertyService.RegisterProperty(ctx, property)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		resp := models.PropertyResponse{PropertyId: propertyId, Message: "Property Created Successfully"}
		response.HandleResponse(w, http.StatusOK, resp)
	}
}

// Reads property details to be updated ad calls update property service to update details
func UpdateProperty(propertyService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		var property models.Property
		err = json.Unmarshal(body, &property)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		getPropertyId := r.URL.Path
		propertyIdPath := strings.Split(getPropertyId, "/")
		propertyId, err := strconv.Atoi(propertyIdPath[2])
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		err = propertyService.UpdateProperty(ctx, property, propertyId)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		response.HandleResponse(w, http.StatusOK, "Property Details Updated Successfully")
	}
}

// Allows to view all properties by calling its service, if the user is a lister it checks if lister wants to see their own properties using owner queryparam. If it is true it allows lister to view their posted properties by calling that service
func GetAllProperties(propertyService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		owner := r.URL.Query().Get("owner")
		role, ok := ctx.Value(constant.RoleKey).(string)
		if !ok {
			response.HandleResponse(w, http.StatusInternalServerError, errhandler.ErrInternalServer.Error())
			return
		}
		var properties []models.Property
		var err error
		if owner == "true" {
			if role == "lister" {
				properties, err = propertyService.GetUsersProperties(ctx)
			}
		} else {
			properties, err = propertyService.GetAllProperties(ctx)
		}

		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		response.HandleResponse(w, http.StatusOK, properties)
	}
}

// Retrieves property id from path and allows to view property of that id by calling get particular property service
func GetParticularProperties(propertyService Service) func(w http.ResponseWriter, r *http.Request) {
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

		properties, err := propertyService.GetParticularProperty(ctx, propertyId)
		if err != nil {
			statusCode, errMessage := errhandler.MapError(err)
			response.HandleResponse(w, statusCode, errMessage)
			return
		}

		response.HandleResponse(w, http.StatusOK, properties)
	}
}
