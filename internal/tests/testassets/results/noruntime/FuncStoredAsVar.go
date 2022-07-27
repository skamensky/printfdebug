package testassest

import "fmt"

var FuncStoredAsVar = func() error {
	fmt.Println("Entering \"FuncStoredAsVar\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"FuncStoredAsVar\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	return nil
}
