package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"se_uf/gator_snapstore/models"

	"gorm.io/gorm"
)

func FetchCartInfo(DB *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var allProductImages []models.ProductCatalogue
	var allImages []models.Image
	allImagesFromDB := DB.Find(&allImages)
	rows, err := allImagesFromDB.Rows()
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()
	var image models.Image
	for rows.Next() {
		DB.ScanRows(rows, &image)
		genres, err := getGenresOfImage(DB, w, image.ImageId)
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		productCatalogueImage := models.ProductCatalogue{
			ImageId:   image.ImageId,
			Price:     image.Price,
			Title:     image.Title,
			WImageURL: image.WImageURL,
			Genre:     genres,
		}
		allProductImages = append(allProductImages, productCatalogueImage)
	}
	SendJSONResponse(w, http.StatusOK, allProductImages)
}

func AddImageToCart(DB *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// TODO: Check if the user is authorized to add to the table or not by comparing the buyerEmailId
	// and the email id from the token received
	type ATCData struct {
		Token        string // TODO: Update this field later accordingly
		BuyerEmailId string
		ImageId      int
	}
	// Parse the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error reading add to cart JSON data ")
		return
	}
	var data ATCData
	err = json.Unmarshal(body, &data)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error unmarshaling ATC data")
		return
	}
	_, flag := checkIfImageExistsOrNot(DB, data.ImageId)
	if !flag {
		SendErrorResponse(w, http.StatusInternalServerError, "Image Id is not valid, adding to cart failed")
		return
	}
	if DB.Create(&models.Cart{
		BuyerEmailId: data.BuyerEmailId,
		ImageId:      data.ImageId,
	}).Error != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error inserting in Cart Schema")
		return
	}
	SendJSONResponse(w, http.StatusOK, map[string]string{"message": "Added to cart"})
}

func DeleteImageFromCart(DB *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// TODO: Check if the user is authorized to delete from the table or not by comparing the buyerEmailId
	// and the email id from the token received
	type DFCData struct {
		Token        string // TODO: Update this field later accordingly
		BuyerEmailId string
		ImageId      int
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error reading delete from cart JSON data ")
		return
	}
	var data DFCData
	err = json.Unmarshal(body, &data)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error unmarshaling DFC data")
		return
	}
	_, flag := checkIfImageExistsOrNot(DB, data.ImageId)
	if !flag {
		SendErrorResponse(w, http.StatusInternalServerError, "Image Id is not valid, deleting from cart failed")
		return
	}
	var dataToBeDeletedFromCart models.Cart
	DB.Where(&models.Cart{BuyerEmailId: data.BuyerEmailId, ImageId: data.ImageId}).Limit(1).Delete(&dataToBeDeletedFromCart)
	SendJSONResponse(w, http.StatusOK, map[string]string{"message": "Removed from cart"})
}
