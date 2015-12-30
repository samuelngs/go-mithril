package mithril

import (
	"fmt"
	"html"
	"regexp"
	"strings"
)

const (
	// Parser is the tags regex parser
	Parser = `(?:(^|#|\.)([^#\.\[\]]+))|(\[.+?\])`
	// AttrParser is the attribute regex parser
	AttrParser = `\[(.+?)(?:=("|'|)(.*?)("|'|))?\]`
)

// M returns VirtualElement
func M(selector string, opts ...interface{}) *VirtualElement {
	// new VirtualElement channel
	c := make(chan *VirtualElement)
	// create element in go routine
	go func() {
		// declare query string
		query := "div"
		// replace query as div if selector does not existed
		if selector != "" {
			query = selector
		}
		// create virtual element
		element := &VirtualElement{"div", []Attribute{}, nil}
		// map options
		for i, opt := range opts {
			switch obj := opt.(type) {
			case []Attribute:
				if i == 0 {
					element.Attrs = obj
				}
			case string, bool, int, int8, int32, int64, uint, uint8, uint32, uint64, float32, float64, complex64, *VirtualElement, *TrustElement, *Component, []interface{}:
				element.Children = obj
			}
		}
		// match all results
		matches := regexp.MustCompile(Parser)
		for _, res := range matches.FindAllStringSubmatch(query, -1) {
			if res[1] == "" && len(res[2]) > 0 {
				element.Tag = strings.ToLower(res[2])
			} else if res[1] == "#" {
				element.Attr("id", res[2])
			} else if res[1] == "." {
				if !element.HasAttr("class") {
					element.Attrs = append(element.Attrs, Class(res[2]))
				} else {
					element.Attr("class", res[2])
				}
			} else if string(res[3][0]) == "[" {
				matches := regexp.MustCompile(AttrParser)
				for _, pair := range matches.FindAllStringSubmatch(res[3], -1) {
					if len(pair[1]) > 0 {
						element.Attrs = append(element.Attrs, NewStringAttr(pair[1], pair[3]))
					}
				}
			}
		}
		c <- element
	}()
	// return virtual element
	return <-c
}

// Trust allows insert HTML for render arbitrary, potentially invalid markup, as well as run arbitrary Javascript, and therefore the developer is responsible for either
func Trust(str string) *TrustElement {
	return &TrustElement{str}
}

// Render returns HTML string
func Render(elements ...interface{}) string {
	str := ""
	for _, element := range elements {
		switch obj := element.(type) {
		case Component:
			obj.Controller()
			str += fmt.Sprintf("%s", obj.View())
		case []interface{}:
			for _, item := range obj {
				str += Render(item)
			}
		case *VirtualElement:
			str += obj.String()
		case *TrustElement:
			str += obj.string
		case string:
			str += html.EscapeString(obj)
		case bool:
			str += fmt.Sprintf("%t", obj)
		case int, int8, int32, int64, uint, uint8, uint32, uint64:
			str += fmt.Sprintf("%d", obj)
		case float32, float64, complex64:
			str += fmt.Sprintf("%f", obj)
		}
	}
	return str
}
