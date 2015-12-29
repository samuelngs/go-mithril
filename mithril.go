package mithril

import (
	"regexp"
	"strings"
)

// M returns VirtualElement
func M(selector string, opts ...interface{}) *VirtualElement {
	// declare query string
	query := "div"
	// replace query as div if selector does not existed
	if selector != "" {
		query = selector
	}
	// create virtual element
	element := &VirtualElement{"div", NewAttributes(), nil}
	// map options
	for i, opt := range opts {
		switch obj := opt.(type) {
		case *Attributes:
			if i == 0 {
				element.Attrs = obj
			}
		case string, bool, int, *VirtualElement:
			element.Children = obj
		}
	}
	// match all results
	re := regexp.MustCompile(`(^[\w\-]+|\#[\w\-]+|\.[\w\-]+|\[[\w\-]+\=\"[\w\-]+\"\]|\[[\w\-]+\=\'[\w\-]+\'\])`)
	for _, res := range re.FindAllStringSubmatch(query, -1) {
		data := res[0]
		value := data[1:len(data)]
		switch string(data[0]) {
		case "[":
			data := strings.Split(value, "=")
			switch len(data) {
			case 1:
				element.Attrs.Data[data[0][:len(data[0])-1]] = ""
			case 2:
				element.Attrs.Data[data[0]] = data[1][1 : len(data[1])-2]
			}
		case "#":
			element.Attrs.ID = value
		case ".":
			element.Attrs.Class = append(element.Attrs.Class, value)
		default:
			element.Tag = data
		}
	}
	// return virtual element
	return element
}
