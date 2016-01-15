package main

import (
	"fmt"

	"github.com/samuelngs/go-mithril"
)

// export global functions
var (
	m      = mithril.M
	render = mithril.Render
	id     = mithril.ID
	class  = mithril.Class
)

// helloComponent
type helloComponent struct {
	*mithril.BaseComponent
}

func (component *helloComponent) Controller() {
	component.Set("sample_text", "Hello World Go-Mithril")
	component.Set("list_items", []string{"object 1", "object 2", "object 3"})
}

func (component *helloComponent) View() interface{} {
	return m("div", []mithril.Attribute{
		id("home"),
		class("main", "main-container"),
	}, []interface{}{
		m("p", component.Get("sample_text")),
		m("ul", func() []interface{} {
			var elements []interface{}
			for _, i := range component.Get("list_items").([]string) {
				elements = append(elements, m("li", i))
			}
			return elements
		}()),
	})
}

func main() {
	component := &helloComponent{mithril.ExtendComponent()}
	html := render(component)
	fmt.Printf("%v\n", html)
}
