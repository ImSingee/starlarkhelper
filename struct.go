package starlarkhelper

import (
	"fmt"
	"go.starlark.net/starlark"
	"strings"
)

type Struct struct {
	Descriptor *StructDescription
	Values     starlark.StringDict
}

var (
	_ starlark.HasAttrs        = (*Struct)(nil)
	_ starlark.Mapping         = (*Struct)(nil)
	_ starlark.Iterable        = (*Struct)(nil)
	_ starlark.IterableMapping = (*Struct)(nil)
	_ Assertable               = (*Struct)(nil)
)

// StructDescription 为一个 struct 的描述信息
type StructDescription struct {
	Name string // struct 名称

	Short      string
	Long       string
	AssertNote string
	Examples   []Example

	Fields map[string]Arg

	Editable                  bool // TODO 是否允许 field 被修改，需要注意的是如果允许那么需要特别处理 assert
	ReturnNoneForUnExistField bool // 对于不存在的 field 返回 None 而不是错误

	StringProvider func(s *Struct) string
	TruthProvider  func(s *Struct) starlark.Bool
	HashProvider   func(s *Struct) (uint32, error)
	AssertProvider func(s *Struct) (pass bool, failReason string)
	GetProvider    func(s *Struct, key starlark.Value) (v starlark.Value, found bool, err error)

	moduleName string // struct 归属的 module 名称
}

// UnsafeConstruct 不做任何检查，创建相应的 struct
func (d *StructDescription) UnsafeConstruct(dict starlark.StringDict) *Struct {
	return &Struct{
		Descriptor: d,
		Values:     dict,
	}
}

func (d *StructDescription) ProvideDoc() DocProvider {
	return NewDocObject(d)
}

// ConstructAll 传递的 dict 必须与期望的 Fields 相对应
// TODO: 类型检测尚未实现
func (d *StructDescription) ConstructAll(dict starlark.StringDict) (*Struct, error) {
	if len(d.Fields) != len(dict) {
		return nil, fmt.Errorf("some fields are mismatch for %s (exepct %d, got %d)", d.Name, len(d.Fields), len(dict))
	}

	for n := range dict {
		if _, ok := d.Fields[n]; !ok {
			return nil, fmt.Errorf("field %s is redundant for %s", n, d.Name)
		}
	}

	return &Struct{
		Descriptor: d,
		Values:     dict,
	}, nil
}

// MustConstructAll ConstructAll but will panic if error
func (d *StructDescription) MustConstructAll(dict starlark.StringDict) *Struct {
	s, err := d.ConstructAll(dict)
	if err != nil {
		panic(err)
	}
	return s
}

func (s *Struct) Description() string {
	buf := new(strings.Builder)

	buf.WriteString(s.Descriptor.Name)
	buf.WriteByte('{')

	first := true
	for k, v := range s.Values {
		if !first {
			buf.WriteString(", ")
		}
		first = false

		buf.WriteString(k)
		buf.WriteString(" = ")
		buf.WriteString(v.String())
	}
	buf.WriteByte('}')

	return buf.String()
}

func (s *Struct) String() string {
	if s.Descriptor.StringProvider == nil {
		return s.Description()
	} else {
		return s.Descriptor.StringProvider(s)
	}
}

func (s *Struct) Type() string { return s.Descriptor.Name }

func (s *Struct) Freeze() {}

func (s *Struct) Truth() starlark.Bool {
	if s.Descriptor.TruthProvider == nil {
		return true
	} else {
		return s.Descriptor.TruthProvider(s)
	}
}

func (s *Struct) Hash() (uint32, error) {
	if s.Descriptor.HashProvider != nil {
		return s.Descriptor.HashProvider(s)
	} else { // copy from starlarkstruct.Struct-Hash
		var x, m uint32 = 8731, 9839
		for k, v := range s.Values {
			namehash, _ := starlark.String(k).Hash()
			x = x ^ 3*namehash
			y, err := v.Hash()
			if err != nil {
				return 0, err
			}
			x = x ^ y*m
			m += 7349
		}
		return x, nil
	}
}

func (s *Struct) Attr(name string) (starlark.Value, error) {
	v, ok := s.Values[name]
	if ok {
		return v, nil
	}

	if s.Descriptor.ReturnNoneForUnExistField {
		return starlark.None, nil
	} else {
		return nil, fmt.Errorf("attr %s is not exist on %s", name, s.Descriptor.Name)
	}
}

func (s *Struct) AttrNames() []string {
	names := make([]string, 0, len(s.Values))
	for k := range s.Values {
		names = append(names, k)
	}
	return names
}

func (s *Struct) Assert() (bool, string) {
	if s.Descriptor.AssertProvider == nil {
		return true, ""
	} else {
		return s.Descriptor.AssertProvider(s)
	}
}

func (s *Struct) Get(key starlark.Value) (v starlark.Value, found bool, err error) {
	if s.Descriptor.GetProvider != nil {
		return s.Descriptor.GetProvider(s, key)
	}

	if k, ok := key.(starlark.String); ok {
		if v, ok := s.Values[string(k)]; ok {
			return v, true, nil
		} else {
			return starlark.None, false, nil
		}
	} else {
		if s.Descriptor.ReturnNoneForUnExistField {
			return starlark.None, false, nil
		} else {
			return nil, false, fmt.Errorf("key must be string")
		}
	}
}

func (s *Struct) Items() []starlark.Tuple {
	items := make([]starlark.Tuple, 0, len(s.Values))

	for k, v := range s.Values {
		items = append(items, starlark.Tuple{starlark.String(k), v})
	}

	return items
}

func (s *Struct) ToDict() *starlark.Dict {
	d := starlark.NewDict(len(s.Values))
	for k, v := range s.Values {
		_ = d.SetKey(starlark.String(k), v)
	}
	return d
}

func (s *Struct) Iterate() starlark.Iterator {
	return s.ToDict().Iterate()
}
