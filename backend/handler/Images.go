package handler

import (
	"bytes"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"se_uf/gator_snapstore/models"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
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

func processWaterMarking(todoImage image.Image, todoImageFormat string) (image.Image, error) {
	image2, err := os.Open("watermark.png")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
		return nil, err
	}
	watermark, err := png.Decode(image2)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
		return nil, err
	}
	defer image2.Close()

	// offset := image.Pt(300, 200)
	// b := todoImage.Bounds()
	baseBound := todoImage.Bounds()
	markBound := watermark.Bounds()
	offset := image.Pt(
		(baseBound.Size().X/2)-(markBound.Size().X/2),
		(baseBound.Size().Y/2)-(markBound.Size().Y/2),
	)
	waterMarkedImage := image.NewRGBA(baseBound)
	draw.Draw(waterMarkedImage, baseBound, todoImage, image.Point{}, draw.Src)
	draw.Draw(waterMarkedImage, watermark.Bounds().Add(offset), watermark, image.Point{}, draw.Over)

	return waterMarkedImage, nil
}

func UploadSellerImage(DB *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// Reading the env file to get the cloud key and secret
	err := godotenv.Load(".env")
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error in reading the env file")
		return
	}
	maxSize := int64(10 * 1024 * 1024) // allow only 10MB of file size
	originalFile, originalFileHeader, err := r.FormFile("sellerImage")
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer originalFile.Close()
	if maxSize < originalFileHeader.Size {
		SendErrorResponse(w, http.StatusInternalServerError, "File too big! Max file size allowed: 10 MB")
		return
	}
	waterMarkedFileHeader := originalFileHeader
	wRawImage, wRawImageformat, err := image.Decode(originalFile)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error in water marking process")
		return
	}
	waterMarkedImage, err := processWaterMarking(wRawImage, wRawImageformat)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error in water marking process")
		return
	}
	watermarkedImageFile, err := os.Create("output" + wRawImageformat)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if wRawImageformat == "jpeg" {
		err := jpeg.Encode(watermarkedImageFile, waterMarkedImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer watermarkedImageFile.Close()
	} else if wRawImageformat == "png" {
		err := png.Encode(watermarkedImageFile, waterMarkedImage)
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	waterMarkedFileHeader.Filename = watermarkedImageFile.Name()
	// waterMarkedFileHeader.Size = watermarkedImageFile.
	// // create an AWS session which can be reused if we're uploading many files
	// s, err := session.NewSession(&aws.Config{
	// 	Region: aws.String("us-east-2"),
	// 	Credentials: credentials.NewStaticCredentials(
	// 		os.Getenv("KEY_ID"), // id
	// 		os.Getenv("SECRET"), // secret
	// 		""),                 // token can be left blank for now
	// })
	// if err != nil {
	// 	SendErrorResponse(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// fileName, err := UploadImageToS3(s, originalFile, originalFileHeader)
	// if err != nil {
	// 	SendErrorResponse(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// fmt.Fprintf(w, "Image uploaded successfully: %v", fileName)

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
		// StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", err
	}
	return tempFileName, err
}

func GetProductInfo(DB *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	imageIdInfo := params["imageId"]
	var image []models.Image
	convertedImageId, err := strconv.Atoi(imageIdInfo)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	currentImageInfo := DB.Where(&models.Image{ImageId: convertedImageId}).Find(&image)
	// For if the user sends the wrong imageId
	if currentImageInfo.RowsAffected <= 0 {
		SendErrorResponse(w, http.StatusInternalServerError, "Invalid Image ID passed")
		return
	}
	imageRow, err := currentImageInfo.Rows()
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer imageRow.Close()
	var currentImage models.Image
	var productImageInfo models.Image
	for imageRow.Next() {
		DB.ScanRows(imageRow, &currentImage)
		productImageInfo = models.Image{
			ImageId:     currentImage.ImageId,
			Title:       currentImage.Title,
			Description: currentImage.Description,
			Price:       currentImage.Price,
			UploadedAt:  currentImage.UploadedAt,
			WImageURL:   currentImage.WImageURL,
		}
	}
	SendJSONResponse(w, http.StatusOK, productImageInfo)
}
