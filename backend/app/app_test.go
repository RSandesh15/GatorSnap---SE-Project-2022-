package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"se_uf/gator_snapstore/models"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetAllImagesAndCompareStruct(t *testing.T) {
	app := initApp()
	image := &models.Image{
		ImageId:   1,
		Title:     "Shooting star",
		Price:     150.25,
		WImageURL: "https://picsum.photos/200",
	}
	// app.DB.Save(image)
	// app.DB.Save(&models.Genre{
	// 	ImageId:   1,
	// 	GenreType: "nature",
	// })
	// Uncomment the following line when running the code for the first time:
	// fillDummyData()
	// app.InsertImage()
	req, _ := http.NewRequest("GET", "/fetchImages", nil)
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(app.getAllImages)

	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusOK, t)
	checkContentType(r, t)
	checkBody(r.Body, image, t)
}

func TestAddToCartWhenExists(t *testing.T) {
    app := initApp()
    var rqBody = toReader(`{"buyerEmailId":"jim@ufl.edu", "ID":10}`)
    req, _ := http.NewRequest("POST", "/addToCart", rqBody)
    r := httptest.NewRecorder()
    handler := http.HandlerFunc(app.addToCart)
    handler.ServeHTTP(r, req)

    checkStatusCode(r.Code, http.StatusOK, t)
    checkContentType(r, t)
    print(r.Body.String())
    type IncomingData struct {
        Message string `json:"message"`
    }
    var dataMap map[string]IncomingData
    err := json.Unmarshal(r.Body.Bytes(), &dataMap)
    if err != nil {
        fmt.Println("Error in Unmarshalling: ", err.Error())
    }
    // processedImageId, _ := strconv.Atoi(imageIdToBePassed)
    if dataMap["data"].Message != "Added to cart" {
        t.Errorf("Add to cart failed")
    }
}

func TestFetchProductInfoWhenExists(t *testing.T) {
	app := initApp()
	req, _ := http.NewRequest("GET", "/fetchProductInfo", nil)
	imageIdToBePassed := "1"
	req = mux.SetURLVars(req, map[string]string{"imageId": imageIdToBePassed})
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(app.getProductInfo)
	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusOK, t)
	checkContentType(r, t)
	// print(r.Body.String())
	var dataMap map[string]models.Image
	err := json.Unmarshal(r.Body.Bytes(), &dataMap)
	if err != nil {
		fmt.Println("Error in Unmarshalling: ", err.Error())
	}
	processedImageId, _ := strconv.Atoi(imageIdToBePassed)
	if dataMap["data"].ImageId != processedImageId {
		t.Errorf("Fetch Product Info does not exist for the given imageId: %v", imageIdToBePassed)
	}
}

func TestFetchProductInfoWhenDoesNotExist(t *testing.T) {
	app := initApp()
	req, _ := http.NewRequest("GET", "/fetchProductInfo", nil)
	imageIdToBePassed := "-123"
	req = mux.SetURLVars(req, map[string]string{"imageId": imageIdToBePassed})
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(app.getProductInfo)
	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusNotFound, t)
	checkContentType(r, t)
	// print(r.Body.String())
	var dataMap map[string]map[string]string
	err := json.Unmarshal(r.Body.Bytes(), &dataMap)
	if err != nil {
		fmt.Println("Error in Unmarshalling: ", err.Error())
	}
	if dataMap["data"]["error"] != "Resource not found" {
		t.Errorf("Fetch Product Info does not exist for the given imageId: %v", imageIdToBePassed)
	}
}

func initApp() App {
	db, _ := gorm.Open(sqlite.Open("gatorsnapstore.db"), &gorm.Config{})
	db.AutoMigrate(&models.Image{})
	db.AutoMigrate(&models.Genre{})
	return App{DB: db}
}

func checkStatusCode(code int, want int, t *testing.T) {
	if code != want {
		t.Errorf("Wrong status code: got %v want %v", code, want)
	}
}

func checkContentType(r *httptest.ResponseRecorder, t *testing.T) {
	ct := r.Header().Get("Content-Type")
	if ct != "application/json" {
		t.Errorf("Wrong Content Type: got %v want application/json", ct)
	}
}

func checkBody(body *bytes.Buffer, image *models.Image, t *testing.T) {
	var dataMap map[string][]models.ProductCatalogue
	err := json.Unmarshal(body.Bytes(), &dataMap)
	// print("Data is: ", data["data"][0].Price)
	if err != nil {
		fmt.Println("Error in Unmarshalling", err.Error())
	}
	// print(body.String())
	// if len(dataMap["data"]) != 1 {
	// 	t.Errorf("Wrong length: got %v want 1", len(dataMap["data"]))
	// }
	firstProductCatalogue := models.ProductCatalogue{
		ImageId:   dataMap["data"][0].ImageId,
		Title:     dataMap["data"][0].Title,
		Price:     float32(dataMap["data"][0].Price),
		Genre:     dataMap["data"][0].Genre,
		WImageURL: dataMap["data"][0].WImageURL,
	}

	// fmt.Println(firstProductCatalogue)
	if image.ImageId != firstProductCatalogue.ImageId {
		t.Errorf("Wrong body: got %v want %v", dataMap["data"][0], image)
	}
}

func toReader(content string) io.Reader {
	return bytes.NewBuffer([]byte(content))
}
