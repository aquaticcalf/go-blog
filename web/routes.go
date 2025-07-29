package web

import (
	"aqclf.xyz/tago"
	"github.com/blog-mini-project/go/pages"
	"log"
	"os"
)

func echoRoutes(app *tago.App) *tago.App {
	app.GET("/", pages.Index())
	return app
}

func buildRoutes() {
	err := os.Mkdir("build", 0755)
	if err != nil {
		log.Fatal("unable to create build folder")
	}
	tago.Build(pages.Index(),"build/index")
}