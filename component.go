package mithril

// Component contains a controller and a view properties.
type Component struct {
	data map[string]interface{}
}

// Controller function creates map for View
func (component *Component) Controller() {
	panic("please override Controller function")
}

// View function creates VirtualElement or string, int, bool
func (component *Component) View() []interface{} {
	panic("please override View function")
	return nil
}
