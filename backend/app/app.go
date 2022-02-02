package app

import (
	"log"
	"net/http"
	"se_uf/gator_snapstore/handler"
	"se_uf/gator_snapstore/models"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) InitializeApplication() {
	db, err := gorm.Open(sqlite.Open("gatorsnapstore.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	a.Router = mux.NewRouter()
	a.DB = db
	a.migrateSchemas()
	a.setRouters()
	// Set the request and response parameters for insertImage()
	// a.insertImage()
	// a.setupGenreCategories()
}

func (a *App) insertImage() {
	// TODO: Read the values from the request parameter r here which is sent from the UI
	for x := 0; x < 20; x++ {
		if a.DB.Create(&models.Image{
			EmailId:     "bruh@ufl.edu",
			Title:       "Shooting star",
			Description: "Good photo!",
			Price:       150.25,
			UploadedAt:  time.Now(),
			ImageURL:    "https://picsum.photos/200", // Insert the original Image url obtained from the bucket
			WImageURL:   "https://picsum.photos/200", // Insert the watermarked Image url obtained from the bucket
		}).Error != nil {
			// handler.SendErrorResponse(w, http.StatusInternalServerError, "Error inserting in Image Schema")
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
		if a.DB.Create(&models.Genre{
			ImageId:   x + 1,
			GenreType: "nature",
			// ImageId: lastInsertedImageId,
		}).Error != nil {
			// handler.SendErrorResponse(w, http.StatusInternalServerError, "Error inserting in Genre Schema")
			return
		}
	}
}

func (a *App) setupGenreCategories() {
	for _, value := range handler.GenreCategorySlice {
		// a.DB.Clauses(clause.Insert{Modifier: "ignore"}).Create(&models.GenreCategories{
		// 	Category: value,
		// })
		if a.DB.Model(&models.GenreCategories{}).Where("category = ?", value).RowsAffected == 0 {
			a.DB.Create(&models.GenreCategories{
				Category: value,
			})
		}
	}
}

func (a *App) migrateSchemas() {
	a.DB.AutoMigrate(&models.Image{})
	a.DB.AutoMigrate(&models.Genre{})
	a.DB.AutoMigrate(&models.GenreCategories{})
}

func (a *App) RunApplication(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func (a *App) setRouters() {
	a.Router.HandleFunc("/fetchImages", a.getAllImages).Methods("GET")
	a.Router.HandleFunc("/fetchGenreCategories", a.getGenreCategories).Methods("GET")
}

func (a *App) getAllImages(w http.ResponseWriter, r *http.Request) {
	handler.GetAllImages(a.DB, w, r)
}

func (a *App) getGenreCategories(w http.ResponseWriter, r *http.Request) {
	handler.GetGenreCategories(a.DB, w, r)
}
