package mithril

import "fmt"

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

// HasAttr returns true if VirtualElement has this attribute in attributes array
func (el *VirtualElement) HasAttr(key string) bool {
	for _, attr := range el.Attrs {
		if attr.Key() == key {
			return true
		}
	}
	return false
}

// IsSelfClosingTag returns true if Tag is a self closing tag element
func (el *VirtualElement) IsSelfClosingTag() bool {
	for _, val := range SelfClosingTagToken {
		if val == el.Tag {
			return true
		}
	}
	return false
}

// String returns the html string of the element
func (el *VirtualElement) String() string {
	res := make(chan string)
	go func() {
		html := "<"
		if el.Tag != "" {
			html += el.Tag
		} else {
			html += "div"
		}
		for _, attr := range el.Attrs {
			html += (" " + attr.Key())
			if len(attr.Val()) > 0 {
				html += ("=\"" + attr.Val() + "\"")
			}
		}
		if el.IsSelfClosingTag() {
			html += " />"
		} else {
			html += ">"
			switch obj := el.Children.(type) {
			case *VirtualElement:
				html += obj.String()
			case string:
				html += obj
			case int, int32, int64:
				html += fmt.Sprintf("%d", obj)
			case float32, float64:
				html += fmt.Sprintf("%f", obj)
			}
			html += "</"
			if el.Tag != "" {
				html += el.Tag
			} else {
				html += "div"
			}
			html += ">"
		}
		res <- html
	}()
	return <-res
}
