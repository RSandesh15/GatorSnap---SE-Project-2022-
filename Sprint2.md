# GatorSnap 
### By: Pulin Soni, Rishab Parmar, Aakansh Togani, Sandesh Ramesh

## Sprint-2 Functionality

### Backend: 


We have worked on and developed the following APIs :

- /uploadSellerImage: When the seller is on the seller upload page, he can upload a new image by entering the details such as title, description, price, genres and the actual image to be uploaded. As soon as the seller hits upload, this API is called, and all the relevant information is stored in the database. Moreover, the process of watermarking on the uploaded image is carried out after calling this API 

- Watermarking: After calling the above-mentioned API, the process of watermarking is applied to the uploaded image. The watermarked image is created by superimposing the watermark on the original image and then manipulating the opacity. Both images are then uploaded to the cloud bucket. Only when the buyer buys the image will the buyer get the unwatermarked image

- Login Authentication and Validation: Using google Oauth we have enable the Login authentication. Google Oauth allows us to obtain an access token in order to be able to access the user's account. The user's ID, emailID and Name is provided by the API endpoint

- /fetchProductInfo: When on the landing page the buyer clicks on any of the images, a new page will open this list all the information of the specific clicked image. This API fetches all the relevant information regarding the clicked image and then provides it to the front end

- /addToCart: We have added the cart functionality in this sprint where the user can add products to the cart. Whenever the user hits add to cart, this API will be called and a mapping between the user and the specific product is added in the database 

- /fetchCartInfo: Whenever the user wishes to see all the products that are there in its cart, this API will be called which provides all the information required 

- /deleteFromCart: Whenever the user wishes to delete any product from its cart, the buyer can do so by hitting the delete from cart button. This will also delete the mapping from the database. So, when the next time the buyer checks out his cart page, the deleted product will not be visible in the cart. 



We have also added the following test cases:

- TestFetchGenreCategories

```
func TestFetchGenreCateogires(t *testing.T) {

	app := initApp()
	app.setupGenreCategories()
	req, _ := http.NewRequest("GET", "/fetchGenreCategories", nil)
	r := httptest.NewRecorder()
	handler_ := http.HandlerFunc(app.getGenreCategories)
	handler_.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusOK, t)
	checkContentType(r, t)

	// print(r.Body.String())

	var dataMap map[string][]string

	err := json.Unmarshal(r.Body.Bytes(), &dataMap)

	if err != nil {
		fmt.Println("Error in Unmarshalling: ", err.Error())
	}
	for index, genre := range handler.GenreCategorySlice {
		if genre != dataMap["data"][index] {
			t.Errorf("Error in fetching genre categories test")
		}
	}
}

```

- TestAddToCartWhenExist

```
func TestAddToCartWhenExists(t *testing.T) {
	app := initApp()
	var rqBody = toReader(`{"buyerEmailId":"jim@ufl.edu", "imageId":10}`)
	req, _ := http.NewRequest("POST", "/addToCart", rqBody)
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(app.addToCart)
	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusOK, t)
	checkContentType(r, t)
	// print(r.Body.String())
	// type IncomingData struct {
	//     Message string `json:"message"`
	// }
	var dataMap map[string]map[string]string
	err := json.Unmarshal(r.Body.Bytes(), &dataMap)
	if err != nil {
		fmt.Println("Error in Unmarshalling: ", err.Error())
	}
	if dataMap["data"]["message"] != "Added to cart" {
		t.Errorf("Add to cart failed when exists")
	}
}
```

- TestAddtoCartWhenDoesNotExist

```
func TestAddToCartWhenDoesNotExists(t *testing.T) {
	app := initApp()
	var rqBody = toReader(`{"buyerEmailId":"jim@ufl.edu", "imageId":-1}`)
	req, _ := http.NewRequest("POST", "/addToCart", rqBody)
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(app.addToCart)
	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusNotFound, t)
	checkContentType(r, t)
	// print(r.Body.String())
	var dataMap map[string]map[string]string
	err := json.Unmarshal(r.Body.Bytes(), &dataMap)
	if err != nil {
		fmt.Println("Error in Unmarshalling: ", err.Error())
	}
	if dataMap["data"]["error"] == "Resource not found" || dataMap["data"]["error"] == "Error unmarshaling" {

	} else {
		t.Errorf("Add to cart when does not exist failed")
	}
}

```

- TestFetchProductInfoWhenExist

```
func TestFetchProductInfoWhenExists(t *testing.T) {
	app := initApp()
	req, _ := http.NewRequest("GET", "/fetchProductInfo", nil)
	imageIdToBePassed := "1"
	req = mux.SetURLVars(req, map[string]string{"imageId": imageIdToBePassed})
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(app.getProductInfo)
	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusOK, t)
	checkContentType(r, t)
	// print(r.Body.String())
	var dataMap map[string]models.Image
	err := json.Unmarshal(r.Body.Bytes(), &dataMap)
	if err != nil {
		fmt.Println("Error in Unmarshalling: ", err.Error())
	}
	processedImageId, _ := strconv.Atoi(imageIdToBePassed)
	if dataMap["data"].ImageId != processedImageId {
		t.Errorf("Fetch Product Info does not exist for the given imageId: %v", imageIdToBePassed)
	}
}

```

- TestFetchProductInfoWhenDoesNotExist

```
func TestFetchProductInfoWhenDoesNotExist(t *testing.T) {
	app := initApp()
	req, _ := http.NewRequest("GET", "/fetchProductInfo", nil)
	imageIdToBePassed := "-123"
	req = mux.SetURLVars(req, map[string]string{"imageId": imageIdToBePassed})
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(app.getProductInfo)
	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusNotFound, t)
	checkContentType(r, t)
	// print(r.Body.String())
	var dataMap map[string]map[string]string
	err := json.Unmarshal(r.Body.Bytes(), &dataMap)
	if err != nil {
		fmt.Println("Error in Unmarshalling: ", err.Error())
	}
	if dataMap["data"]["error"] != "Resource not found" {
		t.Errorf("Fetch Product Info does not exist for the given imageId: %v", imageIdToBePassed)
	}
}

```

- TestGetAllImagesAndCompareStruct 
```
func TestGetAllImagesAndCompareStruct(t *testing.T) {
	app := initApp()
	image := &models.Image{
		ImageId:   1,
		Title:     "Shooting star",
		Price:     150.25,
		WImageURL: "https://picsum.photos/200",
	}
	// app.DB.Save(image)
	// app.DB.Save(&models.Genre{
	// 	ImageId:   1,
	// 	GenreType: "nature",
	// })
	// Uncomment the following line when running the code for the first time:
	// fillDummyData()
	// app.InsertImage()
	req, _ := http.NewRequest("GET", "/fetchImages", nil)
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(app.getAllImages)

	handler.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusOK, t)
	checkContentType(r, t)
	checkBody(r.Body, image, t)
}
``` 


### Frontend:
- / addtoCart : Added functionality of adding to cart features selected by user. Using 'useState' hooks implemented dynamic addition and subtraction of items into the cart.

- / searchBar: created a searchbar for user to find products via title, genre etc. Intend to connect it with backend to display possible matches for the user.

- / Modal : The user can click on details to view a pop-up window. This window calls an API fetchImageDetails. We send imageId to retrieve information of particular image. The data fetched is processed and displayed to the user in this pop-up window. Thus overall making the user interface seamless as the website does not populate all the data in the same webpage but instead offers details in pop-up window.

-/ The User Landing Page is enhanced with Categories of the month, and Popular products section. The webPage uses ample features/components imported via React Material. 

### Video Walkthrough

Here is a walkthrough of what was achieved on the backend and database side for sprint 2. 
<img src='Gifs/PostMan-and-watermarking.gif' title='Backend' width='' />

Here is a walkthrough of what was achieved on the frontend side for sprint 2. 
<img src='Gifs/Frontend(UserLandingPage)-Sprint2.gif' title='Frontend' width='' />
