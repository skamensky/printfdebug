package testassest

import (
	"fmt"
	"log"
)

func FuncLogFatal() {
	fmt.Println("Entering \"FuncLogFatal\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"FuncLogFatal\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	log.Fatal(1)
	fmt.Println("Leaving \"FuncLogFatal\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

func FuncLogFatalln() {
	fmt.Println("Entering \"FuncLogFatalln\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"FuncLogFatalln\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	log.Fatalln(1)

	fmt.Println("Leaving \"FuncLogFatalln\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}
func FuncLogFatalf() {
	fmt.Println("Entering \"FuncLogFatalf\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"FuncLogFatalf\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	log.Fatalf("")

	fmt.Println("Leaving \"FuncLogFatalf\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}
