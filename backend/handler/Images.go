package handler

import (
	"net/http"
	"se_uf/gator_snapstore/models"

	"gorm.io/gorm"
)

func GetAllImages(DB *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// var allProductImages []models.ProductCatalogue
	var allImages []models.Image
	// err := DB.Find(&allProductImages).Error
	err := DB.Find(&allImages).Error
	if(err != nil) {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	// SendJSONResponse(w, http.StatusOK, allProductImages)
	SendJSONResponse(w, http.StatusOK, allImages)
}