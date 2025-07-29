package pages

import (
	"aqclf.xyz/tago"
	"github.com/blog-mini-project/go/layouts"
)

func New() *tago.Element {
	return layouts.DaisyPage(
		tago.Div(
			tago.Class("flex flex-col min-h-[100dvh]"),
			tago.Div(
				tago.Class("w-full max-w-[1300px] mx-auto p-5 grow flex flex-col gap-2"),
				tago.Input( tago.Class("text-3xl focus:outline-none border-none"),
					tago.Attr("placeholder", "title"),
					tago.ID("title"),
				),
				tago.Input( tago.Class("text-lg focus:outline-none border-none"),
					tago.Attr("placeholder", "subtitle"),
					tago.ID("subtitle"),
				),
				tago.Textarea( tago.Class("text-md focus:outline-none border-none grow"),
					tago.Attr("placeholder", "body"),
					tago.ID("body"),
				),
			),
		),
	)
}
