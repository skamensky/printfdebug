package testassest

import "fmt"

func FuncPanic() {
	fmt.Println("Entering \"FuncPanic\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"FuncPanic\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	panic("Oh no, how will we know how we got here?")
	fmt.Println("Leaving \"FuncPanic\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	return
}
