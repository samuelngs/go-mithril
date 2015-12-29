package mithril

// ExtendComponent creates a basic component view
func ExtendComponent() *Component {
	return &Component{make(map[string]interface{})}
}

// Component contains a controller and a view properties.
type Component struct {
	controller map[string]interface{}
}

// Controller function creates map for View
func (component *Component) Controller() {
	panic("please override Controller function")
}

// View function creates VirtualElement or string, int, bool
func (component *Component) View() interface{} {
	panic("please override View function")
	return nil
}

// Get returns controller value
func (component *Component) Get(key string) interface{} {
	return component.controller[key]
}

// Set value to controller scope
func (component *Component) Set(key string, val interface{}) {
	component.controller[key] = val
}
