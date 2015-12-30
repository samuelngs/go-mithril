# go-mithril

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
