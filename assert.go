package starlarkhelper

// Assertable 为数据提供 assert 的能力
type Assertable interface {
	Assert() (pass bool, failReason string)
}

// AssertFunc TODO
func AssertFunc() {

}
