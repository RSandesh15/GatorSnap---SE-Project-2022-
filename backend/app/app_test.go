package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"se_uf/gator_snapstore/models"
	"testing"
	"time"

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
	// fillDummyData()
	req, _ := http.NewRequest("GET", "/fetchImages", nil)
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(app.getAllImages)

	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusOK, t)
	checkContentType(r, t)
	checkBody(r.Body, image, t)
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
		fmt.Println("=================================>", err.Error())
	}
	// print(body.String())
	if len(dataMap["data"]) != 1 {
		t.Errorf("Wrong length: got %v want 1", len(dataMap["data"]))
	}
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

func fillDummyData() {
	app := initApp()
	for x := 0; x < 20; x++ {
		// TODO: Reading from the request parameter r for finding the corresponding values
		if app.DB.Create(&models.Image{
			SellerEmailId: "bruh@ufl.edu",
			Title:         "Shooting star",
			Description:   "Good photo!",
			Price:         150.25,
			UploadedAt:    time.Now(),
			ImageURL:      "https://picsum.photos/200", // Insert the original Image url obtained from the bucket
			WImageURL:     "https://picsum.photos/200", // Insert the watermarked Image url obtained from the bucket
		}).Error != nil {
			// handler.SendErrorResponse(w, http.StatusInternalServerError, "Error inserting in Image Schema")
			fmt.Printf("Error inserting in Image Schema")
		}
		// var lastImage models.Image
		// temp := a.DB.Last(&models.Image)
		// row, err  := temp.Rows()
		// if err != nil {
		// 	handler.SendErrorResponse(w, http.StatusInternalServerError, "Error inserting in Genre Schema")
		// 	return
		// }
		// a.DB.ScanRows(row, lastImage)
		// lastInsertedImageId := lastImage.ImageId
		// // Loop for all the available genres passed from the front end
		if app.DB.Create(&models.Genre{
			ImageId:   x + 1,
			GenreType: "nature",
			// ImageId: lastInsertedImageId,
		}).Error != nil {
			// handler.SendErrorResponse(w, http.StatusInternalServerError, "Error inserting in Genre Schema")
			return
		}
	}
}
