package web

import (
	"aqclf.xyz/tago"
	"log"
)

func App() {
	app := tago.NewApp()
	err := echoRoutes(app).Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
