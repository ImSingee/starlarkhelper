package starlarkhelper

import (
	"fmt"
	"go.starlark.net/starlark"
)

// ToStringDict convert elements to StringDict
// any element must be one of
// - *StarlarkModule
// - *Module
// - starlark.Value
// Panic: any other values will result in panic
func ToStringDict(globalMiddleware Middleware, elements ...map[string]interface{}) starlark.StringDict {
	d := starlark.StringDict{}

	for _, element := range elements {
		for k, e := range element {
			switch e := e.(type) {
			case *Function:
				e = e.Copy()
				e.Middleware = ChainMiddleware(globalMiddleware, e.Middleware)
				d[k] = e
			case *Module:
				e = e.Copy()
				e.FuncMiddleware = ChainMiddleware(globalMiddleware, e.FuncMiddleware)
				d[k] = e.Get()
			case starlark.Value:
				d[k] = e
			default:
				panic(fmt.Sprintf("%T", e))
			}
		}
	}

	return d
}
