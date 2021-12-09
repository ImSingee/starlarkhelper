package starlarkhelper

import (
	"fmt"
	"go.starlark.net/starlark"
)

func TraverseLists(values []starlark.Value, f func(starlark.Value) error) error {
	for _, value := range values {
		err := TraverseList(value, f)
		if err != nil {
			return err
		}
	}

	return nil
}

func TraverseList(value starlark.Value, f func(starlark.Value) error) error {
	if value == nil {
		return nil
	}

	if list, ok := value.(*starlark.List); ok {
		for i := 0; i < list.Len(); i++ {
			err := f(list.Index(i))
			if err != nil {
				return err
			}
		}

		return nil
	} else {
		return f(value)
	}
}

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

func NewListFromStrings(elems []string) *starlark.List {
	values := make([]starlark.Value, len(elems))

	for i, e := range elems {
		values[i] = starlark.String(e)
	}

	return starlark.NewList(values)
}
