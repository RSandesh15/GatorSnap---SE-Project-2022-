package handler

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"se_uf/gator_snapstore/models"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func FetchCartInfo(DB *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// TODO: Check if the user is authorized to add to the table or not by comparing the buyerEmailId
	// and the email id from the token received
	params := mux.Vars(r)
	buyerEmailId := params["buyerEmailId"]
	allCartProducts, err := fetchCartRecords(DB, w, buyerEmailId)
	if err != nil {
		return
	}
	SendJSONResponse(w, http.StatusOK, allCartProducts)
}

func fetchCartRecords(DB *gorm.DB, w http.ResponseWriter, buyerEmailId string) ([]models.ProductCatalogue, error) {
	var buyerCartProducts models.Cart
	allBuyerCartProducts := DB.Where(&models.Cart{BuyerEmailId: buyerEmailId}).Find(&buyerCartProducts)
	rows, err := allBuyerCartProducts.Rows()
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return nil, err
	}
	var allProductImages []models.ProductCatalogue
	defer rows.Close()
	var cartInfo models.Cart
	for rows.Next() {
		DB.ScanRows(rows, &cartInfo)
		image, flag := checkIfImageExistsOrNot(DB, cartInfo.ImageId)
		if !flag {
			SendErrorResponse(w, http.StatusInternalServerError, "Mentioned imageId does not exist in fetch cart info method")
			return nil, errors.New("custom error")
		}
		imageRow, err := image.Rows()
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
			return nil, err
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
	return allProductImages, nil
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
		SendErrorResponse(w, http.StatusNotFound, "Error unmarshaling")
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
		SendErrorResponse(w, http.StatusNotFound, "Error unmarshaling")
		return
	}
	_, flag := checkIfImageExistsOrNot(DB, data.ImageId)
	if !flag {
		SendErrorResponse(w, http.StatusNotFound, "Resource not found")
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

func CheckoutAndProcessPayment(DB *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// TODO: Check if the user is authorized to delete from the table or not by comparing the buyerEmailId
	// and the email id from the token received
	// First computing the total sum of the amount:
	type CAPPData struct {
		Token        string // TODO: Update this field later accordingly
		BuyerEmailId string
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error reading Checkout and process payment from cart JSON data ")
		return
	}
	var data CAPPData
	err = json.Unmarshal(body, &data)
	if err != nil {
		SendErrorResponse(w, http.StatusNotFound, "Error unmarshaling")
		return
	}
	allCartProducts, err := fetchCartRecords(DB, w, data.BuyerEmailId)
	if err != nil {
		return
	}
	var amount int64 = 0.0
	for _, cartProduct := range allCartProducts {
		amount = amount + int64(cartProduct.Price)
	}
	amount = amount * 100
	// Processing the payment using Stripe payment gateway:
	err = godotenv.Load(".env")
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error in reading the env file")
		return
	}
	// Loading the stripe secret key from the environment variables
	stripe.Key = os.Getenv("STRIPE_SECRET")
	params := &stripe.PaymentIntentParams{
		Amount:       stripe.Int64(amount),
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		ReceiptEmail: &data.BuyerEmailId,
	}
	params.AddMetadata("buyerEmailId", data.BuyerEmailId)
	pi, err := paymentintent.New(params)

	// println("pi.New: %v", pi.ClientSecret)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("pi.New: %v", err)
		return
	}
	// The output of this API call is that a client secret is returned to the front end with all the information needed
	// to make the payment to the user. Following this, the same client_secret will be returned from the frontend in the
	// API '/processPayment'
	SendJSONResponse(w, http.StatusOK, map[string]string{"clientSecret": pi.ClientSecret})
}

func EmailProduct(DB *gorm.DB, w http.ResponseWriter, r *http.Request) {
	type EPData struct {
		Token           string // TODO: Update this field later accordingly
		BuyerEmailId    string
		PaymentIntentId string
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error reading email product JSON data ")
		return
	}
	var data EPData
	err = json.Unmarshal(body, &data)
	if err != nil {
		SendErrorResponse(w, http.StatusNotFound, "Error unmarshaling")
		return
	}

	// Fetching the payment data using paymentIntentId:
	err = godotenv.Load(".env")
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error in reading the env file")
		return
	}
	stripe.Key = os.Getenv("STRIPE_SECRET")
	processedPaymentIntent, err := paymentintent.Get(data.PaymentIntentId, nil)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error in extracting data from payment intent")
		return
	}
	paymentMetadata := processedPaymentIntent.Metadata
	buyerEmailId := paymentMetadata["buyerEmailId"]
	// println("Email id from the metadata", buyerEmailId, processedPaymentIntent.Amount)

	// TODO: Check if the email id from the token, paymentIntent and email address sent are all the same, only then proceed

	// Fetching all the cart details from the clients secret:
	allCartProducts, err := fetchCartRecords(DB, w, buyerEmailId)
	if err != nil {
		return
	}
	// Mailing every product in the buyers cart and updating the previous orders table at the same time
	for _, cartProduct := range allCartProducts {
		// Fetching the imageData that gives us info about the seller email id, the original image url and the amount
		imageData, err := fetchSingleProduct(DB, w, cartProduct.ImageId)
		if err != nil {
			return
		}

		fileName := cartProduct.Title + ".jpeg"
		err = downloadOriginalImageFromCloud(imageData.ImageURL, fileName)
		if err != nil {
			println(err)
			SendErrorResponse(w, http.StatusInternalServerError, "Error downloading image from cloud in email product API")
			return
		}

		// Sending email with attachment (original unwatermarked image) to the buyer:
		msg := gomail.NewMessage()
		msg.SetHeader("From", os.Getenv("GATORSSUSERNAME"))
		msg.SetHeader("To", buyerEmailId)
		msg.SetHeader("Subject", "Your Gator SnapStore purchase is here! ID: "+cartProduct.Title+"_"+strconv.Itoa(rand.Int()))
		msg.SetBody("text/html", "Thanks for doing business with Gator SnapStore. We have attached your image below. Enjoy it! Go Gators! \nRegards,\nGator SnapStore.")
		msg.Attach(fileName)

		n := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("GATORSSUSERNAME"), os.Getenv("GATORSSPASS"))
		// Send the email
		if err := n.DialAndSend(msg); err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Deleting the original Image that has been downloaded after a mail has been sent to the user and the previous orders
		// table has been correctly updated:
		err = os.Remove(fileName)
		if err != nil {
			print("Error in deleting originalImage.jpeg file: ", err.Error())
		}

		// Updating the buyer cart state
		var dataToBeDeletedFromCart models.Cart
		rowsDeleted := DB.Where(&models.Cart{BuyerEmailId: buyerEmailId, ImageId: cartProduct.ImageId}).Delete(&dataToBeDeletedFromCart)
		// As the delete query is not running correctly with limit 1, we are deleting the record from the database and then adding the records
		// such that addedRecords = deleteRecords - 1
		for i := 1; i < int(rowsDeleted.RowsAffected); i++ {
			if DB.Create(&models.Cart{
				BuyerEmailId: buyerEmailId,
				ImageId:      cartProduct.ImageId,
			}).Error != nil {
				SendErrorResponse(w, http.StatusInternalServerError, "Error inserting in Cart Schema in deletion operation")
				return
			}
		}

		// Inserting data in the previous orders table:
		if DB.Create(&models.PreviousOrders{
			BuyerEmailId: buyerEmailId,
			SellerEmailId: imageData.SellerEmailId,
			ImageId: imageData.ImageId,
			Title: imageData.Title,
			Price: imageData.Price,
			BoughtAt: time.Time{},
		}).Error != nil {
			SendErrorResponse(w, http.StatusInternalServerError, "Error inserting in Previous Orders Schema")
			return
		}
		
		// TODO: remove the following break statement (for testing)
		break
	}
	SendJSONResponse(w, http.StatusOK, map[string]string{"message": "Order placed and shipped successfully! Your order has been delivered to your registered email id!"})
}

func downloadOriginalImageFromCloud(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}
