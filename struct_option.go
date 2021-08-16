package starlarkhelper

func NewStruct(name string, args ...Arg) *StructDescription {
	m := make(map[string]Arg, len(args))
	for _, a := range args {
		m[a.InName()] = a
	}

	return &StructDescription{
		Name:   name,
		Fields: m,
	}
}

func (d *StructDescription) WithAssertProvider(f func(s *Struct) (pass bool, failReason string)) *StructDescription {
	d.AssertProvider = f
	return d
}
