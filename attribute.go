package mithril

import "strings"

// Attribute is an attribute namespace-key-value triple.
type Attribute interface {
	Key(...string) string
	Val(...string) string
	Namespace(...string) string
	Add(...string) string
	Remove(...string) string
}

// NewStringAttr Creates a StringAttribute object
func NewStringAttr(args ...string) *StringAttribute {
	attr := &StringAttribute{}
	for i, arg := range args {
		switch i {
		case 0:
			attr.key = arg
		case 1:
			attr.val = arg
		case 2:
			attr.namespace = arg
		}
	}
	return attr
}

// StringAttribute is an attribute namespace-key-value triple.
type StringAttribute struct {
	namespace, key string
	val            string
}

// Key for sets or returns key string
func (attr *StringAttribute) Key(keys ...string) string {
	for _, k := range keys {
		attr.key = k
		break
	}
	return attr.key
}

// Val for sets or returns attr value
func (attr *StringAttribute) Val(vals ...string) string {
	for _, v := range vals {
		attr.val = v
		break
	}
	return attr.val
}

// Namespace for sets and returns namespace value
func (attr *StringAttribute) Namespace(namespaces ...string) string {
	for _, n := range namespaces {
		attr.namespace = n
		break
	}
	return attr.namespace
}

// Add for adding value(s) to attr store (Do nothing for StringAttribute)
func (attr *StringAttribute) Add(vals ...string) string {
	return attr.val
}

// Remove for removing value(s) from attr value variable (Do nothing for StringAttribute)
func (attr *StringAttribute) Remove(vals ...string) string {
	return attr.val
}

// Separator for change specific text separator when join string array (Do nothing for StringAttribute)
func (attr *StringAttribute) Separator(s string) *StringAttribute {
	return attr
}

// NewListAttr Creates a ListAttribute object
func NewListAttr(args ...string) *ListAttribute {
	attr := &ListAttribute{val: []string{}}
	for i, arg := range args {
		if i == 0 {
			attr.key = arg
		} else {
			attr.val = append(attr.val, arg)
		}
	}
	return attr
}

// NewNSListAttr creates a ListAttribute object
func NewNSListAttr(args ...string) *ListAttribute {
	attr := &ListAttribute{val: []string{}}
	for i, arg := range args {
		if i == 0 {
			attr.key = arg
		} else if i < len(args)-1 {
			attr.val = append(attr.val, arg)
		} else {
			attr.namespace = arg
		}
	}
	return attr
}

// ListAttribute is an attribute namespace-key-value triple.
type ListAttribute struct {
	namespace, key, separator string
	val                       []string
}

// Key for sets or returns key string
func (attr *ListAttribute) Key(keys ...string) string {
	for _, k := range keys {
		attr.key = k
		break
	}
	return attr.key
}

// Val for sets or returns attr value
func (attr *ListAttribute) Val(vals ...string) string {
	for _, v := range vals {
		attr.Add(v)
	}
	return strings.Join(attr.val[:], attr.separator)
}

// Contains return the result of if attr contains a specific value
func (attr *ListAttribute) Contains(v string) bool {
	for _, b := range attr.val {
		if b == v {
			return true
		}
	}
	return false
}

// Namespace for sets and returns namespace value
func (attr *ListAttribute) Namespace(namespaces ...string) string {
	for _, n := range namespaces {
		attr.namespace = n
		break
	}
	return attr.namespace
}

// Add for adding value(s) to attr store
func (attr *ListAttribute) Add(vals ...string) string {
	for _, v := range vals {
		if !attr.Contains(v) {
			attr.val = append(attr.val, v)
		}
	}
	return strings.Join(attr.val[:], attr.separator)
}

// Remove for removing value(s) from attr value variable
func (attr *ListAttribute) Remove(vals ...string) string {
	for _, v := range vals {
		for i, e := range attr.val {
			if v == e {
				attr.val = append(attr.val[:i], attr.val[i+1:]...)
			}
		}
	}
	return strings.Join(attr.val[:], attr.separator)
}

// Separator for change specific text separator when join string array
func (attr *ListAttribute) Separator(s string) Attribute {
	attr.separator = s
	return attr
}

// NewIDAttr returns new class attribute
func NewIDAttr(vals ...string) Attribute {
	attr := NewStringAttr("id")
	for _, v := range vals {
		attr.Add(v)
	}
	return attr
}

// NewClassAttr returns new class attribute
func NewClassAttr(vals ...string) Attribute {
	attr := NewListAttr("class").Separator(" ")
	for _, v := range vals {
		attr.Add(v)
	}
	return attr
}

// NewStyleAttr returns new class attribute
func NewStyleAttr(vals ...string) Attribute {
	attr := NewListAttr("style").Separator(";")
	for _, v := range vals {
		attr.Add(v)
	}
	return attr
}
