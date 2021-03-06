package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"crypto/rand"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	//"gorm.io/gorm"
)

type googleAuthResponse struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

var (
	state = "holderState"
)

type userClaims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

/*
func Register(DB *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var data map[string]string //we nedd to declare a new User struct

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}*/
//+ os.Getenv("PORT")
func oAuthGoogleConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL: "http://localhost:8085/google/callback",
		ClientID:    "305686927939-hsc849g4qd7jtuqbepl2dlf58or3p42l.apps.googleusercontent.com",
		//ClientID:  os.Getenv("GOOGLECLIENTID"),
		//ClientSecret: os.Getenv("GOOGLECLIENSECRET"),
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func GenerateRandomString() (string, error) {
	n := 5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	s := fmt.Sprintf("%X", b)
	return s, nil
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {

	//tempState, err := GenerateRandomString()
	//state := tempState
	//if err != nil {
	//	SendErrorResponse(w, http.StatusInternalServerError, "Error in randomizing string")
	//	return
	//}
	url := oAuthGoogleConfig().AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != state {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	token, err := oAuthGoogleConfig().Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Error in Token callback")
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Could not create request")
		return
	}
	defer resp.Body.Close()
	//googleResponse, err:= ioutil.ReadAll(resp.Body)

	googleResponse := googleAuthResponse{}
	err = json.NewDecoder(resp.Body).Decode(&googleResponse)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Could not parse response")
		return
	}
	fmt.Println(googleResponse.Email)
	//fmt.Fprint(w, "Response: %s", googleResponse)
	//err=json.Unmarshal(googleResponse, )

	tkn, _ := GenerateToken(&googleResponse)
	fmt.Println(tkn)
	//cookie, err := r.Cookie("insomnia")
	cookie := http.Cookie{
		Name:     "insomnia",
		Value:    tkn,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: false,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
	fmt.Println(cookie.Name)
	http.Redirect(w, r, "http://localhost:3000/userLandingPage", http.StatusTemporaryRedirect)
}

func GenerateToken(googleResponse *googleAuthResponse) (string, error) {

	claims := userClaims{
		Email: googleResponse.Email,
		//Name:  googleResponse.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    googleResponse.ID,
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := t.SignedString([]byte("JWT_SECRET"))

	return tokenString, err
}

func ValidateToken(signedToken string) (claims *userClaims, err error) {

	token, err := jwt.ParseWithClaims(signedToken, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("JWT_SECRET"), nil ///jwt_secret???????
	})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok {
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return
	}

	return claims, err
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "insomnia",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
	fmt.Println(cookie.Name)
	http.Redirect(w, r, "http://localhost:3000/login", http.StatusTemporaryRedirect)
}

/*
callback function to be changed
func GoogleCallback() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.FormValue("state") != state {
			return c.Redirect("/", http.StatusTemporaryRedirect)
		}
		token, err := oAuthGoogleConfig().Exchange(context.Background(), c.FormValue("code"))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.SendString("Error in Token Callback")
		}

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			return c.SendString("Cannot get your details bro")
		}
		defer resp.Body.Close()
		googleResponse := googleAuthResponse{}
		err = json.NewDecoder(resp.Body).Decode(&googleResponse)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON("Error")
		}

		user := models.User{
			Id:    primitive.NewObjectID(),
			Name:  googleResponse.Name,
			Email: googleResponse.Email,
		}
		userCollection := database.MI.Db.Collection("users")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		findResult := userCollection.FindOne(ctx, bson.M{
			"email": user.Email,
		})

		if err := findResult.Err(); err != nil {
			_, err := userCollection.InsertOne(ctx, user)
			if err != nil {
				return err
			}
		}

		tkn, _ := GenerateToken(&googleResponse)
		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    tkn,
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		}
		c.Cookie(&cookie)

		return c.Redirect("http://localhost:"+os.Getenv("CLIENT_PORT")+"/home", http.StatusTemporaryRedirect)
	}
}
*/
