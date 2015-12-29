package mithril

// VirtualElement is an object which can convert to real DOM element
type VirtualElement struct {
	Tag      string
	Attrs    *Attributes
	Children interface{}
}
