package main

import (
	"log"
	"net/http"
	"se_uf/gator_snapstore/app"

	"github.com/rs/cors"
)

func main() {
	app := &app.App{}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	handler := c.Handler(app.Router)
	log.Fatal(http.ListenAndServe(":3000", handler))

	app.InitializeApplication()
	app.RunApplication(":8085")
}
