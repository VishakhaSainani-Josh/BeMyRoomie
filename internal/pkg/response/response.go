package response

import (
	"encoding/json"

	"net/http"
)

func HandleResponse(w http.ResponseWriter, statusCode int, message any) {
	if err, ok := message.(error); ok {
		message = err.Error()
	}
	response, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Error converting response to json", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)

}
