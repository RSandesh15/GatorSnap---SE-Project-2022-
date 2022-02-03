package handler

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"se_uf/gator_snapstore/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/globalsign/mgo/bson"
	"github.com/joho/godotenv"
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

func getGenresOfImage(DB *gorm.DB, w http.ResponseWriter, imageId int) ([]string, error) {
	var genres []string
	var genre []models.Genre
	// Another way of finding data:
	// genresOfCurrentImage := DB.Where("image_id = ?", strconv.Itoa(imageId)).Find(&genre)
	genresOfCurrentImage := DB.Where(&models.Genre{ImageId: imageId}).Find(&genre)
	genreRows, err := genresOfCurrentImage.Rows()
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return genres, err
	}
	defer genreRows.Close()
	var currentGenre models.Genre
	for genreRows.Next() {
		DB.ScanRows(genreRows, &currentGenre)
		genres = append(genres, currentGenre.GenreType)
	}
	return genres, nil
}

func GetGenreCategories(DB *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var genreCategories []models.GenreCategories
	allCategoriesFromDB := DB.Find(&genreCategories)
	genreCategoriesRows, err := allCategoriesFromDB.Rows()
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer genreCategoriesRows.Close()
	var allGenreCategories []string
	var genreCategory models.GenreCategories
	for genreCategoriesRows.Next() {
		DB.ScanRows(genreCategoriesRows, &genreCategory)
		allGenreCategories = append(allGenreCategories, genreCategory.Category)
	}
	SendJSONResponse(w, http.StatusOK, allGenreCategories)
}

func UploadSellerImage(DB *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// Reading the env file to get the cloud key and secret
	err := godotenv.Load(".env")
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error in reading the env file")
	}
	maxSize := int64(10 * 1024 * 1024) // allow only 10MB of file size
	file, fileHeader, err := r.FormFile("sellerImage")
	// wRawImage, format, err := image.Decode(file)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if maxSize < fileHeader.Size {
		SendErrorResponse(w, http.StatusInternalServerError, "File too big! Max file size allowed: 10 MB")
		return
	}
	defer file.Close()
	// create an AWS session which can be reused if we're uploading many files
	s, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("KEY_ID"), // id
			os.Getenv("SECRET"), // secret
			""),                 // token can be left blank for now
	})
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	fileName, err := UploadImageToS3(s, file, fileHeader)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Fprintf(w, "Image uploaded successfully: %v", fileName)

	// Inserting the image details in the database:

	// Inserting the genre details in the database:
}

// UploadFileToS3 saves a file to aws bucket and returns the url to // the file and an error if there's any
func UploadImageToS3(s *session.Session, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	// get the file size and read
	// the file content into a buffer
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	// create a unique file name for the file
	tempFileName := "pictures/" + bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)

	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String("gator-snapstore"),
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", err
	}
	return tempFileName, err
}
