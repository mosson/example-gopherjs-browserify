package jslist

import (
	"fmt"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

// List is list item
type List struct {
	Text string
	Href string
}

// New returns new List
func New(text string, href string) *js.Object {
	return js.MakeWrapper(&List{Text: text, Href: href})
}

// Node returns DOM Node
func (l *List) Node() *js.Object {
	document := js.Global.Get("document")
	object := document.Call("createElement", "p")
	object.Set("innerHTML", fmt.Sprintf("<a href=\"%s\">%s</a>", l.Href, l.Text))
	return object
}

// Routine returns bool through goroutine
func (l *List) Routine() *js.Object {
	ch := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- true
	}()

	return js.MakeWrapper(<-ch)
}
