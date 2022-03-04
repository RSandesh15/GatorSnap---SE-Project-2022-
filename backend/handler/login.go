package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

type googleAuthResponse struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
}

type userClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func Register(DB *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var data map[string]string //we nedd to declare a new User struct

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func oAuthGoogleConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  "http://localhost:" + os.Getenv("PORT") + "/google/callback",
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func GenerateToken(googleResponse *googleAuthResponse) (string, error) {

	claims := userClaims{
		Email: googleResponse.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    googleResponse.ID,
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := t.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return tokenString, err
}

func ValidateToken(signedToken string) (claims *userClaims, err error) {

	token, err := jwt.ParseWithClaims(signedToken, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
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
