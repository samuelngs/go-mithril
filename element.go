package mithril

// VirtualElement is an object which can convert to real DOM element
type VirtualElement struct {
	Tag      string
	Attrs    *Attributes
	Children interface{}
}

// ID returns the attribute ID of the VirtualElement instance
func (el *VirtualElement) ID() string {
	return el.Attrs.ID
}
