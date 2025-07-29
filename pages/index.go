package pages

import (
	"aqclf.xyz/tago"
	"github.com/blog-mini-project/go/layouts"
)

func Index() *tago.Element {
	return layouts.DaisyPage(
		tago.Div(
			tago.Class("flex flex-col min-h-[100dvh]"),
		),
	)
}