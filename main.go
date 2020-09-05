package main

import (
	"api-ajf/app"
	"log"
)

func main() {
	appx := app.New()
	app.Route(appx)
	err := appx.Listen("127.0.0.1:3000")
	if err != nil {
		log.Fatal(err)
	}
}
