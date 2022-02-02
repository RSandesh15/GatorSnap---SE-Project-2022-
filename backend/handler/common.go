package handler

import (
	"encoding/json"
	"net/http"
)

var GenreCategorySlice = [...]string{"nature", "space", "abstract", "silhouette", "adventure", "architecture", "sunsets"}

func SendJSONResponse(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(status)
	w.Write(response)
}

func SendErrorResponse(w http.ResponseWriter, errorCode int, errorMessage string) {
	SendJSONResponse(w, errorCode, map[string]string{"error": errorMessage})
}
