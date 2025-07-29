package layouts

import (
	. "aqclf.xyz/tago"
	"github.com/blog-mini-project/go/config"
	"github.com/blog-mini-project/go/components"
)

func DaisyPage(children *Element) *Element {
	return Html(
		Attr("data-theme", config.Config.Theme),
		Head(
			Daisy(),
			Title(config.Config.Title),
			Link(
				Attr("rel", "shortcut icon"),
				Attr("href", config.Config.Favicon),
				Attr("type", "image/png"),
			),
			Meta(
				Attr("name", "viewport"),
				Attr("content", "width=device-width, initial-scale=1.0"),
			),
		),
		Body(
			components.Nav(),
			children,
		),
	)
}