package starlarkhelper

type symbolArg string

func (n symbolArg) isArg()                        {}
func (n symbolArg) isArgs()                       {}
func (n symbolArg) InName() string                { return string(n) }
func (n symbolArg) InType() string                { return "" }
func (n symbolArg) InDefaultValue() string        { return "" }
func (n symbolArg) OutType() string               { return "" }
func (n symbolArg) Optional() Arg                 { return n }
func (n symbolArg) Default(value interface{}) Arg { return n }
func (n symbolArg) Short(s string) Arg            { return n }
func (n symbolArg) Long(s string) Arg             { return n }
func (n symbolArg) DocShort() string              { return "" }
func (n symbolArg) DocLong() string               { return "" }

var STAR = symbolArg("*")
var SLASH = symbolArg("/")
var ELLIPSIS = symbolArg("...")
