package main

import (
	"github.com/blog-mini-project/go/config"
	"github.com/blog-mini-project/go/web"
)


func main() {
	config.Load()

	if config.Config.Mode == "dev" {
		web.App()
	} else {
		web.Build()
	}
}