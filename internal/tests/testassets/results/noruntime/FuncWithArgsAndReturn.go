package testassest

import "fmt"

func FuncWithArgsAndReturn(s2 string) error {
	fmt.Println("Entering \"FuncWithArgsAndReturn\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"FuncWithArgsAndReturn\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	// this is a comment on its own line
	/*
		This is a multi line
		comment
	*/
	return nil //This is inline comment
}
