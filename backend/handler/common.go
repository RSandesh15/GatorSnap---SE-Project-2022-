package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"se_uf/gator_snapstore/models"
)

var GenreCategorySlice = [...]string{"nature", "space", "abstract", "silhouette", "adventure", "architecture", "sunsets"}

func SendJSONResponse(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(map[string]interface{}{"data": payload})
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

//will have to call this function once the data from the form is converted and sent in json format to server
func DecodeJson() {
	jsonDataSellerForm := []byte(`
	{
		"EmailId": "thishouldwork@ufl.edu",
		"ImageId": "2",
		"Title": "Shooting Star",
		"Description": "Depicts the very famous....",
		"Price": "15.99",
		"UploadedAt": "time",
		"ImageURL": "IURL:1b3a9374-1a8e-434e-90ab-21aa7b9b80e7"
		"WImageURL": "WURL:1b3a9374-1a8e-434e-90ab-21aa7b9b80e7"
	  }
	  `)

	var checkjson models.Image

	checkValid := json.Valid(jsonDataSellerForm)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataSellerForm, &checkjson)
		//how to upload valid json in DB??
		fmt.Printf("%#v\n", checkjson)
	} else {
		fmt.Println("JSON was not valid")
	}
}
