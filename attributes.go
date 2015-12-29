package mithril

// NewAttributes creates new attribute instance
func NewAttributes() *Attributes {
	return &Attributes{
		ID:    "",
		Class: []string{},
		Style: make(map[string]interface{}),
		Data:  make(map[string]string),
		Name:  "",
		Type:  "",
		Href:  "",
		Value: "",
	}
}

// Attributes represents DOM attributes
type Attributes struct {
	ID    string
	Class []string
	Style map[string]interface{}
	Data  map[string]string
	Name  string
	Type  string
	Href  string
	Value string
}
