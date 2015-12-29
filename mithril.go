package mithril

import "regexp"

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
		matches := regexp.MustCompile(`(?:(^|#|\.)([^#\.\[\]]+))|(\[.+?\])`)
		for _, res := range matches.FindAllStringSubmatch(query, -1) {
			if res[1] == "" && len(res[2]) > 0 {
				element.Tag = res[2]
			} else if res[1] == "#" {
				element.Attrs.ID = res[2]
			} else if res[1] == "." {
				element.Attrs.Class = append(element.Attrs.Class, res[2])
			} else if string(res[3][0]) == "[" {
				matches := regexp.MustCompile(`\[(.+?)(?:=("|'|)(.*?)("|'|))?\]`)
				for _, pair := range matches.FindAllStringSubmatch(res[3], -1) {
					if len(pair[1]) > 0 {
						element.Attrs.Data[pair[1]] = pair[3]
					}
				}
			}
		}
		c <- element
	}()
	// return virtual element
	return <-c
}
