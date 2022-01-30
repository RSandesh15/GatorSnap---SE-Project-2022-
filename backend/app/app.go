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
	a.insertDummyData()
}

func (a *App) insertDummyData() {
	//Format MM-DD-YYYY
	for x := 0; x < 20; x++ {
		a.DB.Create(&models.Image{
			Title:       "Shooting star",
			EmailId:     "bruh@ufl.edu", // get this from the front end, or update it to be better such that it uses JWT
			ImageId:     x,              // Do this bit dynamically, use uuid or some similar library to do this
			Description: "Good photo!",
			Price:       150.25,
			UploadedAt:  time.Now(),
			ImageURL:    "https://picsum.photos/200", // Set after uploading to Firebase/ S3
			WImageURL:   "https://picsum.photos/200", // Set after uploading to Firebase/ S3
		})
		println("sdfsfsd")
	}
	println("Yolo!")
}

func (a *App) migrateSchemas() {
	a.DB.AutoMigrate(&models.Image{})
	a.DB.AutoMigrate(&models.Genre{})
}

func (a *App) RunApplication(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func (a *App) setRouters() {
	a.Router.HandleFunc("/fetchImages", a.getAllImages).Methods("GET")
}

func (a *App) getAllImages(w http.ResponseWriter, r *http.Request) {
	handler.GetAllImages(a.DB, w, r)
}
