package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"status":      "success",
			"status_code": http.StatusOK,
			"message":     data["message"],
			"data":        data["data"],
		})

}

//Respond Not Found 404
func RespondNotFound(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"status":      "error",
			"status_code": http.StatusNotFound,
			"message":     data["message"],
			"data":        nil,
		})
}

func RespondBadRequest(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"status":      "error",
			"status_code": http.StatusBadRequest,
			"message":     data["message"],
			"data":        nil,
		})
}

func RespondMethodNotAllowed(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"status":      "error",
			"status_code": http.StatusMethodNotAllowed,
			"message":     data["message"],
			"data":        nil,
		})
}
