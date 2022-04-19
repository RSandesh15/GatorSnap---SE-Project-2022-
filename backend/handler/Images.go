package handler

import (
	"bytes"
	"errors"
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
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
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

func processWaterMarking(todoImage image.Image) (image.Image, error) {
	image2, err := os.Open("watermark_gator.png")
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
	// Setting up the AWS S3 bucket
	s, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("KEY_ID"),
			os.Getenv("SECRET"),
			""),
	})
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Processing the images:
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

	// Uploading the original image to AWS S3 bucket:
	originalImageURL, err := UploadImageToS3(s, originalFile, originalFileHeader)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	println(w, "Image uploaded successfully: %v", originalImageURL)

	// Watermarking the image to be uploaded:
	// Resetting the seek to 0, 0 after reading the original file
	originalFile.Seek(0, 0)
	wRawImage, wRawImageformat, err := image.Decode(originalFile)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error in water marking process decoding")
		return
	}
	waterMarkedImage, err := processWaterMarking(wRawImage)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error in water marking process")
		return
	}
	watermarkedImageFile, err := os.Create("output." + wRawImageformat)
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

	// Uploading the waterMarked image to AWS S3 bucket:
	// Reading the contents of the waterMarkedImageFile to obtain its information:
	file, err := os.Open(watermarkedImageFile.Name())
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	// Making a buffer according to the size of the newly created water marked image
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// create a unique file name for the file
	tempFileName := "pictures/" + bson.NewObjectId().Hex()
	waterMarkedImageURL, err := s3ConfigUpload(s, tempFileName, size, buffer)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	println(w, "WaterMarked Image uploaded successfully: %v", waterMarkedImageURL)
	err = os.Remove("output.jpeg")
	if err != nil {
		print("Error in deleting output.jpeg file: ", err.Error())
	}

	// Parsing the multipart file data:
	err = r.ParseMultipartForm(0)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error parsing multipart file data")
		return
	}
	sellerEmailId := r.FormValue("sellerEmailId")
	title := r.FormValue("title")
	description := r.FormValue("description")
	priceInString := r.FormValue("price")
	price, err := strconv.ParseFloat(priceInString, 32)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error parsing multipart file data")
		return
	}
	// Inserting the image details in the database:
	if DB.Create(&models.Image{
		SellerEmailId: sellerEmailId,
		Title:         title,
		Description:   description,
		Price:         float32(price),
		UploadedAt:    time.Now(),
		ImageURL:      originalImageURL,
		WImageURL:     waterMarkedImageURL,
	}).Error != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error inserting in Image Schema")
		return
	}
	var genres []string
	r.ParseForm()
	for key, value := range r.Form {
		if key == "genres[]" {
			genres = value
			break
		}
	}
	var allImages *models.Image
	var lastInsertedImageId int64
	DB.Find(&allImages).Count(&lastInsertedImageId)
	for _, genre := range genres {
		if DB.Create(&models.Genre{
			ImageId:   int(lastInsertedImageId),
			GenreType: string(genre),
		}).Error != nil {
			SendErrorResponse(w, http.StatusInternalServerError, "Error inserting in Genre Schema")
			return
		}
	}
	SendJSONResponse(w, http.StatusOK, map[string]string{"message": "Image uploaded successfully"})
}

// UploadFileToS3 saves a file to aws bucket and returns the url to // the file and an error if there's any
func UploadImageToS3(s *session.Session, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	// get the file size and read
	// the file content into a buffer
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	tempFileName := "pictures/" + bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)

	fileURL, err := s3ConfigUpload(s, tempFileName, size, buffer)
	if err != nil {
		return "", err
	}
	return fileURL, nil
}

func s3ConfigUpload(s *session.Session, tempFileName string, size int64, buffer []byte) (string, error) {
	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket: aws.String("gator-snapstore"),
		Key:    aws.String(tempFileName),
		// ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
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
	fileURL := "https://gator-snapstore.s3.amazonaws.com/" + tempFileName
	return fileURL, err
}

func GetProductInfo(DB *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	imageIdInfo := params["imageId"]
	convertedImageId, err := strconv.Atoi(imageIdInfo)
	if err != nil {
		SendErrorResponse(w, http.StatusNotFound, "Resource not found")
		return
	}
	// TODO: Handle the condition when an alphabet is passed as imageId to the API instead of an integer
	productImageInfo, err := fetchSingleProduct(DB, w, convertedImageId)
	if err != nil {
		return
	}
	SendJSONResponse(w, http.StatusOK, productImageInfo)
}

func fetchSingleProduct(DB *gorm.DB, w http.ResponseWriter, convertedImageId int) (models.Image, error) {
	var productImageInfo models.Image
	currentImageInfo, flag := checkIfImageExistsOrNot(DB, convertedImageId)
	if !flag {
		SendErrorResponse(w, http.StatusNotFound, "Resource not found")
		return productImageInfo, errors.New("custom error")
	}
	imageRow, err := currentImageInfo.Rows()
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return productImageInfo, err
	}
	defer imageRow.Close()
	var currentImage models.Image
	for imageRow.Next() {
		DB.ScanRows(imageRow, &currentImage)
		productImageInfo = models.Image{
			ImageId:       currentImage.ImageId,
			SellerEmailId: currentImage.SellerEmailId,
			Title:         currentImage.Title,
			Description:   currentImage.Description,
			Price:         currentImage.Price,
			UploadedAt:    currentImage.UploadedAt,
			WImageURL:     currentImage.WImageURL,
			ImageURL:      currentImage.ImageURL,
		}
	}
	return productImageInfo, nil
}

func checkIfImageExistsOrNot(DB *gorm.DB, imageId int) (*gorm.DB, bool) {
	var image []models.Image
	currentImageInfo := DB.Where(&models.Image{ImageId: imageId}).Find(&image)
	if currentImageInfo.RowsAffected <= 0 {
		return nil, false
	}
	return currentImageInfo, true
}
