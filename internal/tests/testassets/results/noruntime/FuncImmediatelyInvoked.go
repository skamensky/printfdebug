package testassest

import "fmt"

var FuncImmediatelyInvoked = func() error {
	fmt.Println("Entering \"FuncImmediatelyInvoked\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"FuncImmediatelyInvoked\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	return nil
}()
