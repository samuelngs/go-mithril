package mithril

import "testing"

// NewHelloComponent creates a hello component view
func NewHelloComponent() *HelloComponent {
	return &HelloComponent{ExtendComponent()}
}

type HelloComponent struct {
	*BaseComponent
}

func (component *HelloComponent) Controller() {
	component.Set("sample_text", "Hello World Go-Mithril")
	component.Set("list_items", []string{"object 1", "object 2", "object 3"})
}

func (component *HelloComponent) View() interface{} {
	return M("div", []Attribute{
		NewIDAttr("home"),
		NewClassAttr("main", "main-container"),
	}, []interface{}{
		M("p", component.Get("sample_text")),
		M("ul", func() []interface{} {
			var elements []interface{}
			for _, i := range component.Get("list_items").([]string) {
				elements = append(elements, M("li", i))
			}
			return elements
		}()),
	})
}

func TestHelloComponent(t *testing.T) {
	component := NewHelloComponent()
	component.Controller()
	el := component.View()
	expected := "<div id class=\"main main-container\"><p>Hello World Go-Mithril</p><ul><li>object 1</li><li>object 2</li><li>object 3</li></ul></div>"
	if el == expected {
		t.Fatalf("Expected element to be a %s object but it was %s", expected, el)
	}
}
