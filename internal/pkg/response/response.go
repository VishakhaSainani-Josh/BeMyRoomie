package response

import (
	"encoding/json"

	"net/http"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/models"
)

func HandleResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := models.HttpResponse{
		Message: message,
		Data:    data,
	}

	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error converting response to json", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(marshaledResponse)
}
