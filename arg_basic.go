package starlarkhelper

// basicArg 基本参数
type basicArg struct {
	name         string
	type_        BasicType
	optional     bool        // 可选
	defaultValue interface{} // 默认值
}

func (b *basicArg) isArg()                    {}
func (b *basicArg) isArgs()                   {}
func (b *basicArg) InName() string            { return b.name }
func (b *basicArg) Optional() Arg             { b.optional = true; return b }
func (b *basicArg) Default(v interface{}) Arg { b.defaultValue = v; return b }
func (b *basicArg) Short(s string) Arg        { return withShort(b, s) }
func (b *basicArg) Long(l string) Arg         { return withLong(b, l) }
func (b *basicArg) DocShort() string          { return "" }
func (b *basicArg) DocLong() string           { return "" }

func (b *basicArg) InType() string {
	if b.optional {
		return string(b.type_) + "?"
	} else {
		return string(b.type_)
	}
}

func (b *basicArg) InDefaultValue() string { return defaultValueToString(b.defaultValue) }

func (b *basicArg) OutType() string {
	if b.optional {
		return "<" + string(b.type_) + " / None>"
	} else {
		return string(b.type_)
	}
}

func IntArg(name string) Arg {
	return &basicArg{name: name, type_: TypeInt}
}

func FloatArg(name string) Arg {
	return &basicArg{name: name, type_: TypeFloat}
}

func BoolArg(name string) Arg {
	return &basicArg{name: name, type_: TypeBool}
}

func StringArg(name string) Arg {
	return &basicArg{name: name, type_: TypeString}
}

func BytesArg(name string) Arg {
	return &basicArg{name: name, type_: TypeBytes}
}

func AnyArg(name string) Arg {
	return &basicArg{name: name, type_: TypeAny}
}
