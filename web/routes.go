package web

import (
	"aqclf.xyz/tago"
	"github.com/blog-mini-project/go/pages"
	"log"
	"os"
)

func echoRoutes(app *tago.App) *tago.App {
	app.GET("/", pages.Index())
	app.GET("/new", pages.New())
	return app
}

func buildRoutes() {
	buildDir := "build"

	if _, err := os.Stat(buildDir); os.IsNotExist(err) {
		err := os.Mkdir(buildDir, 0755)
		if err != nil {
			log.Fatal("unable to create build folder")
		}
	}
	tago.Build(pages.Index(), "build/index")
}
