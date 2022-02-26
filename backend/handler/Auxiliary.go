package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"se_uf/gator_snapstore/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func FetchCartInfo(DB *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// TODO: Check if the user is authorized to add to the table or not by comparing the buyerEmailId
	// and the email id from the token received
	var allProductImages []models.ProductCatalogue
	params := mux.Vars(r)
	buyerEmailId := params["buyerEmailId"]
	var buyerCartProducts models.Cart
	allBuyerCartProducts := DB.Where(&models.Cart{BuyerEmailId: buyerEmailId}).Find(&buyerCartProducts)
	rows, err := allBuyerCartProducts.Rows()
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()
	var cartInfo models.Cart
	for rows.Next() {
		DB.ScanRows(rows, &cartInfo)
		image, flag := checkIfImageExistsOrNot(DB, cartInfo.ImageId)
		if !flag {
			SendErrorResponse(w, http.StatusInternalServerError, "Mentioned imageId does not exist in fetch cart info method")
			return
		}
		imageRow, err := image.Rows()
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer imageRow.Close()
		var currentImage models.Image
		for imageRow.Next() {
			DB.ScanRows(imageRow, &currentImage)
			productCatalogueImage := models.ProductCatalogue{
				ImageId:   currentImage.ImageId,
				Price:     currentImage.Price,
				Title:     currentImage.Title,
				WImageURL: currentImage.WImageURL,
			}
			allProductImages = append(allProductImages, productCatalogueImage)
		}
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
		SendErrorResponse(w, http.StatusNotFound, "Error unmarshaling ATC data")
		return
	}
	// TODO: Check if the user buyerEmailId exists or not
	// _, flag := checkIfBuyerEmailIdExistsOrNot(DB, data.ImageId)
	// if !flag {
	// 	SendErrorResponse(w, http.StatusNotFound, "Image Id is not valid, adding to cart failed")
	// 	return
	// }
	_, flag := checkIfImageExistsOrNot(DB, data.ImageId)
	if !flag {
		SendErrorResponse(w, http.StatusNotFound, "Resource not found")
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
		SendErrorResponse(w, http.StatusNotFound, "Error unmarshaling DFC data")
		return
	}
	_, flag := checkIfImageExistsOrNot(DB, data.ImageId)
	if !flag {
		SendErrorResponse(w, http.StatusInternalServerError, "Resource not found")
		return
	}
	var dataToBeDeletedFromCart models.Cart
	rowsDeleted := DB.Where(&models.Cart{BuyerEmailId: data.BuyerEmailId, ImageId: data.ImageId}).Delete(&dataToBeDeletedFromCart)
	// As the delete query is not running correctly with limit 1, we are deleting the record from the database and then adding the records
	// such that addedRecords = deleteRecords - 1
	for i := 1; i < int(rowsDeleted.RowsAffected); i++ {
		if DB.Create(&models.Cart{
			BuyerEmailId: data.BuyerEmailId,
			ImageId:      data.ImageId,
		}).Error != nil {
			SendErrorResponse(w, http.StatusInternalServerError, "Error inserting in Cart Schema in deletion operation")
			return
		}
	}
	SendJSONResponse(w, http.StatusOK, map[string]string{"message": "Removed from cart"})
}
