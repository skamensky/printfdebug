package testassest

import "fmt"

func (s *S) StructMethodWithReturn() error {
	fmt.Println("Entering \"StructMethodWithReturn\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"StructMethodWithReturn\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	return nil
}
