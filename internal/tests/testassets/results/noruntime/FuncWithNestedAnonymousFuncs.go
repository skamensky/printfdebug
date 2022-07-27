package testassest

import "fmt"

func FuncWithNestedAnonymousFuncs() {
	fmt.Println("Entering \"FuncWithNestedAnonymousFuncs\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	innerFunc := func() {
		fmt.Println("Entering \"innerFunc\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		nestedInnerFunc := func() {
			fmt.Println("Entering \"nestedInnerFunc\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
			fmt.Println("Leaving \"nestedInnerFunc\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		}
		nestedInnerFunc()
		fmt.Println("Leaving \"innerFunc\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	}
	_ = innerFunc
	fmt.Println("Leaving \"FuncWithNestedAnonymousFuncs\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}
