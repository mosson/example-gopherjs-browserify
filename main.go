package main

import (
	"js/jslist"
	"net/url"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Module.Get("exports").Set("list", map[string]interface{}{
		"New":   jslist.New,
		"Parse": url.Parse,
	})
}
