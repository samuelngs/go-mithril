package mithril

var (
	// SelfClosingTagToken is a list of self closing elements tag
	SelfClosingTagToken = []string{"area", "base", "br", "col", "command", "embed", "hr", "img", "input", "keygen", "link", "meta", "param", "source", "track", "wbr", "!doctype"}
)

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

// String returns the html string of the element
func (el *VirtualElement) String() string {
	return ""
}
