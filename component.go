package mithril

// Component for Mithril
type Component interface {
	Controller()
	View() interface{}
}

// ExtendComponent creates a basic component view
func ExtendComponent() *BaseComponent {
	return &BaseComponent{make(map[string]interface{})}
}

// BaseComponent contains a controller and a view properties.
type BaseComponent struct {
	scope map[string]interface{}
}

// Controller function creates map for View
func (component *BaseComponent) Controller() {
	panic("please override Controller function")
}

// View function creates VirtualElement or string, int, bool
func (component *BaseComponent) View() interface{} {
	panic("please override View function")
	return nil
}

// Get returns controller value
func (component *BaseComponent) Get(key string) interface{} {
	return component.scope[key]
}

// Set value to controller scope
func (component *BaseComponent) Set(key string, val interface{}) {
	component.scope[key] = val
}
