package components

import (
	. "aqclf.xyz/tago"
	. "github.com/blog-mini-project/go/config"
)

func Nav() *Element {
	return Div(
		Class("navbar bg-base-100 shadow-sm"),
		A(
			Class("btn btn-ghost text-xl"),
			Config.Title,
			Attr("href", Config.BaseUrl),
		),
		Div(Class("grow")),
		newPost(),
	)
}


func newPost() *Element {
	if Config.Mode == "dev" {
		return A(
			Class("btn btn-ghost text-xl"),
			"new post",
			Attr("href", Config.BaseUrl + "/new"),
		)
	}
	return Fragment()
}