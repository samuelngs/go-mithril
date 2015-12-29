package mithril

var (
	// SelfClosingTagToken is a list of self closing elements tag
	SelfClosingTagToken = []string{"area", "base", "br", "col", "command", "embed", "hr", "img", "input", "keygen", "link", "meta", "param", "source", "track", "wbr", "!doctype"}
)

// VirtualElement is an object which can convert to real DOM element
type VirtualElement struct {
	Tag      string
	Attrs    []Attribute
	Children interface{}
}

// Attr returns or sets the attribute
func (el *VirtualElement) Attr(key string, val ...string) string {
	for _, attr := range el.Attrs {
		if attr.Key() == key {
			if len(val) > 0 {
				attr.Val(val[0])
			}
			return attr.Val()
		}
	}
	if len(val) > 0 {
		el.Attrs = append(el.Attrs, NewStringAttr(key, val[0]))
		return val[0]
	}
	return ""
}

// String returns the html string of the element
func (el *VirtualElement) String() string {
	return ""
}
