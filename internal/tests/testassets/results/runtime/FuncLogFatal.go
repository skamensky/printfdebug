package testassest

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func FuncLogFatal() {
	printfdebug_Printf_FuncLogFatal("Entering \"FuncLogFatal\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	printfdebug_Printf_FuncLogFatal("Leaving \"FuncLogFatal\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	log.Fatal(1)
	printfdebug_Printf_FuncLogFatal("Leaving \"FuncLogFatal\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

func FuncLogFatalln() {
	printfdebug_Printf_FuncLogFatal("Entering \"FuncLogFatalln\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	printfdebug_Printf_FuncLogFatal("Leaving \"FuncLogFatalln\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	log.Fatalln(1)

	printfdebug_Printf_FuncLogFatal("Leaving \"FuncLogFatalln\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}
func FuncLogFatalf() {
	printfdebug_Printf_FuncLogFatal("Entering \"FuncLogFatalf\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	printfdebug_Printf_FuncLogFatal("Leaving \"FuncLogFatalf\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	log.Fatalf("")

	printfdebug_Printf_FuncLogFatal("Leaving \"FuncLogFatalf\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

var _ = runtime.Caller   // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = filepath.Clean   // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = fmt.Println      // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = strings.Split    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = os.PathSeparator // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_FuncLogFatal(message string, pathDepthFromEnd int) {
	minInt := func(first int, second int) (min int) {
		if first < second {
			return first
		} else {
			return second
		}
	}

	_, file, line, ok := runtime.Caller(1)
	if ok {
		fileParts := strings.Split(file, string(os.PathSeparator))
		pathFromEndSafe := minInt(len(fileParts), pathDepthFromEnd)
		limited := filepath.Join(fileParts[len(fileParts)-pathFromEndSafe:]...)
		limitedCleaned := "??"
		if limited != "" {
			limitedCleaned = limited
		}
		fmt.Printf("%v:%v %v\n", limitedCleaned, line, message)
	} else {
		fmt.Printf("unkown_file:? %v\n", message)
	}
}
