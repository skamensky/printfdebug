package testassest

import (
	"fmt"
	"os"
)

func FuncOsExit() {
	fmt.Println("Entering \"FuncOsExit\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"FuncOsExit\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	os.Exit(1)
	fmt.Println("Leaving \"FuncOsExit\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}
