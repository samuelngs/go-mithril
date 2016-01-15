# go-mithril

[![Build Status](https://travis-ci.org/samuelngs/go-mithril.svg?branch=master)](https://travis-ci.org/samuelngs/go-mithril)
[![Coverage Status](https://coveralls.io/repos/samuelngs/go-mithril/badge.svg?branch=master&service=github)](https://coveralls.io/github/samuelngs/go-mithril?branch=master)
[![GoDoc](https://godoc.org/github.com/samuelngs/go-mithril?status.svg)](https://godoc.org/github.com/samuelngs/go-mithril)

Mithril (Isomorphic) server-side render for Go

### What is Mithril(.JS)?

[Mithril.js](https://github.com/lhorie/mithril.js) is a small and fast MVC JavaScript framework. It encourages an architecture similar to Angular.js, and uses a virtual DOM like React.js, all while avoiding the need for libraries like jQuery. Mithril's small size and API makes it ideal for embedded JavaScript widgets and user interfaces that have high performance requirements.

Github: [https://github.com/lhorie/mithril.js](https://github.com/lhorie/mithril.js)

### What is Go-Mithril?

Go-Mithril is a way to create reusable Mithril UI components for both server-side and client-side.

### Isomorphic means?

i·so·mor·phic: “corresponding or similar in form and relations”

## Usage

To create an username textbox:
```go
M("input.username[type='text']") // Returns <input class="username" type="text" />
```
Create a textbox with id, class and type
```go
M("input", []Attribute{
	ID("username"),
	Class("input", "input-username"),
	NewStringAttr("type", "text"),
}) // Returns <input id="username" class="input input-username" type="text" />
```
Create a div with content
```go
M("div", "Hello World") // Returns <div>Hello World</div>
```

Create a list with component
```go
type ListComponent struct {
	*BaseComponent
}

func (component *ListComponent) Controller() {
	component.Set("list_items", []string{"object 1", "object 2", "object 3"})
}

func (component *ListComponent) View() interface{} {
	return M("ul", func() []interface{} {
        var elements []interface{}
        for _, i := range component.Get("list_items").([]string) {
            elements = append(elements, M("li", i))
        }
        return elements
    }()),
}

func main() {
  mithril.Render(&ListComponent{ExtendComponent()}) // Returns <ul><li>object 1</li><li>object 2</li><li>object 3</li></ul>
}
```

## Shortcut

To use `m` instead of `M`
```go
import (
	"github.com/samuelngs/go-mithril"
)
var (
	m      = mithril.M
	render = mithril.Render
	id     = mithril.ID
	class  = mithril.Class
)
func main() {
 	render(
 		m("div", "Hello World")
 	)
}
```

## Documentation

`go doc` format documentation for this project can be viewed online without installing the package by using the GoDoc page at: https://godoc.org/github.com/samuelngs/go-mithril

## Contributing

Everyone is encouraged to help improve this project. Here are a few ways you can help:

- [Report bugs](https://github.com/samuelngs/go-mithril/issues)
- Fix bugs and [submit pull requests](https://github.com/samuelngs/go-mithril/pulls)
- Write, clarify, or fix documentation
- Suggest or add new features

## License

This project is distributed under the MIT license found in the [LICENSE](./LICENSE) file.

```
The MIT License (MIT)

Copyright (c) 2015 Samuel

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
