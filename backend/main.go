package main

import "se_uf/gator_snapstore/app"

func main() {
	app := &app.App{}

	app.InitializeApplication()
	app.RunApplication(":8888")
}