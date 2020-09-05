package main

import (
	"api-ajf/app"
	"log"
)

var appx *app.AppX

func main() {
	appx = app.New()
	appx.Route()
	err := appx.Listen()
	if err != nil {
		log.Fatal(err)
	}

}
