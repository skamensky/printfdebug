package testassest

import "fmt"

func (s *S) StructMethodWithArgsAndReturn(s2 string) error {
	fmt.Println("Entering \"StructMethodWithArgsAndReturn\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"StructMethodWithArgsAndReturn\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	return nil
}
