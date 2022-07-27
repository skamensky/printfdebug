package testassest

import "fmt"

func AnonymousFunc() {
	fmt.Println("Entering \"AnonymousFunc\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	func() {
		fmt.Println("Entering \"anonymous-function Fun\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		fmt.Println("Leaving \"anonymous-function Fun\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	}()
	fmt.Println("Leaving \"AnonymousFunc\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}
