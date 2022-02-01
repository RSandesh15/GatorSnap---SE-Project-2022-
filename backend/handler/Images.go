package handler

import (
	"net/http"
	"se_uf/gator_snapstore/models"

	"gorm.io/gorm"
)

func GetAllImages(DB *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var allProductImages []models.ProductCatalogue
	var allImages []models.Image
	allImagesFromDB := DB.Find(&allImages)
	rows, err := allImagesFromDB.Rows()
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
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

func getGenresOfImage(DB *gorm.DB, w http.ResponseWriter, imageId int) ([]string, error) {
	var genres []string
	var allGenres []models.Genre
	err := DB.Find(&allGenres).Where("ImageId = ?", imageId).Select("GenreType").Error
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return genres, err
	}
	return genres, nil
}
