package app

import (
	"fmt"
	"log"
	"net/http"
	"se_uf/gator_snapstore/handler"
	"se_uf/gator_snapstore/models"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	// a.InsertImage()
	a.setupGenreCategories()
}

func (a *App) InsertImage() {
	// Read the values from the request parameter r here which is sent from the UI
	for x := 0; x < 20; x++ {
		// TODO: Reading from the request parameter r for finding the corresponding values
		if a.DB.Create(&models.Image{
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
		if a.DB.Create(&models.Genre{
			ImageId:   x + 1,
			GenreType: "nature",
		}).Error != nil {
			return
		}
	}
}

func (a *App) setupGenreCategories() {
	for _, value := range handler.GenreCategorySlice {
		a.DB.Clauses(clause.Insert{Modifier: "or ignore"}).Create(&models.GenreCategories{
			Category: value,
		})
	}
}

func (a *App) migrateSchemas() {
	a.DB.AutoMigrate(&models.Image{})
	a.DB.AutoMigrate(&models.Genre{})
	a.DB.AutoMigrate(&models.GenreCategories{})
	a.DB.AutoMigrate(&models.Cart{})
	a.DB.AutoMigrate(&models.PreviousOrders{})
}

func (a *App) RunApplication(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func (a *App) setRouters() {
	a.Router.HandleFunc("/fetchImages", a.getAllImages).Methods("GET")
	//a.Router.HandleFunc("/postform", postFormHandler).Methods("POST")
	a.Router.HandleFunc("/fetchGenreCategories", a.getGenreCategories).Methods("GET")
	a.Router.HandleFunc("/uploadSellerImage", a.uploadSellerImage).Methods("POST")
	a.Router.HandleFunc("/fetchProductInfo/{imageId}", a.getProductInfo).Methods("GET")
	a.Router.HandleFunc("/fetchCartInfo/{buyerEmailId}", a.fetchCartInfo).Methods("GET")
	a.Router.HandleFunc("/addToCart", a.addToCart).Methods("POST")
	a.Router.HandleFunc("/deleteFromCart", a.deleteFromCart).Methods("POST")
	a.Router.HandleFunc("/checkoutAndProcessPayment", a.checkoutAndProcessPayment).Methods("POST")
	a.Router.HandleFunc("/emailProduct", a.emailProduct).Methods("POST")
	a.Router.HandleFunc("/fetchSellerTransactions/{sellerEmailId}", a.fetchSellerTransactions).Methods("GET")
	a.Router.HandleFunc("/google/login", a.googleLogin).Methods("GET")
	//one to handle callback
}

func (a *App) getAllImages(w http.ResponseWriter, r *http.Request) {
	handler.GetAllImages(a.DB, w, r)
}

// var tpl *template.Template
//tpl, _ = tpl.ParseGlob("templates/*.html")
// template is the folder where the sellerpageform is stored

//func (a *App) postFormHandler(w http.ResponseWriter, r *http.Request) {
//tpl.ExecuteTemplate(w, "postform.html", nil)
//}
func (a *App) getGenreCategories(w http.ResponseWriter, r *http.Request) {
	handler.GetGenreCategories(a.DB, w, r)
}

func (a *App) uploadSellerImage(w http.ResponseWriter, r *http.Request) {
	handler.UploadSellerImage(a.DB, w, r)
}

func (a *App) getProductInfo(w http.ResponseWriter, r *http.Request) {
	handler.GetProductInfo(a.DB, w, r)
}

func (a *App) fetchCartInfo(w http.ResponseWriter, r *http.Request) {
	handler.FetchCartInfo(a.DB, w, r)
}

func (a *App) addToCart(w http.ResponseWriter, r *http.Request) {
	handler.AddImageToCart(a.DB, w, r)
}

func (a *App) deleteFromCart(w http.ResponseWriter, r *http.Request) {
	handler.DeleteImageFromCart(a.DB, w, r)
}

func (a *App) checkoutAndProcessPayment(w http.ResponseWriter, r *http.Request) {
	handler.CheckoutAndProcessPayment(a.DB, w, r)
}

func (a *App) emailProduct(w http.ResponseWriter, r *http.Request) {
	handler.EmailProduct(a.DB, w, r)
}

func (a *App) fetchSellerTransactions(w http.ResponseWriter, r *http.Request) {
	handler.FetchSellerTransactions(a.DB, w, r)
}

func (a *App) googleLogin(w http.ResponseWriter, r *http.Request) {
	handler.GoogleLogin(w, r)
}
