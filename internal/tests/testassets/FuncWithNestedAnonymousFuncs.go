package testassest

func FuncWithNestedAnonymousFuncs() {
	innerFunc := func() {
		nestedInnerFunc := func() {}
		nestedInnerFunc()
	}
	_ = innerFunc
}
