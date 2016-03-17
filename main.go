package main

import (
	"js/jslist"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Module.Get("exports").Set("list", map[string]interface{}{
		"New": jslist.New,
	})
}
