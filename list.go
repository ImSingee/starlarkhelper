package starlarkhelper

import (
	"fmt"
	"go.starlark.net/starlark"
)

func ToStringList(value starlark.Value) ([]string, error) {
	if value == nil {
		return nil, nil
	}

	switch v := value.(type) {
	case starlark.String:
		return []string{v.GoString()}, nil
	case starlark.Iterable:
		it := v.Iterate()
		defer it.Done()

		result := make([]string, 0, 16)

		var vv starlark.Value
		for it.Next(&vv) {
			vvv, ok := vv.(starlark.String)
			if !ok {
				return nil, fmt.Errorf("list item must be string")
			}

			result = append(result, vvv.GoString())
		}

		return result, nil
	}

	return nil, fmt.Errorf("invalid type %T", value)
}
