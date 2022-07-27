package testassest

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func FuncWithExplicitReturn() {
	printfdebug_Printf_FuncWithExplicitReturn("Entering \"FuncWithExplicitReturn\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	printfdebug_Printf_FuncWithExplicitReturn("Leaving \"FuncWithExplicitReturn\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	return
}

var _ = runtime.Caller // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = filepath.Clean // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = fmt.Println    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_FuncWithExplicitReturn(message string, pathDepthFromEnd int) {
	maxInt := func(first int, second int) (max int) {
		if first > second {
			return first
		} else {
			return second
		}
	}

	_, file, line, ok := runtime.Caller(1)
	if ok {
		fileParts := filepath.SplitList(file)
		pathFromEndSafe := maxInt(len(fileParts), pathDepthFromEnd)
		limited := filepath.Join(fileParts[pathFromEndSafe:]...)
		limitedCleaned := "??"
		if limited != "" {
			limitedCleaned = limited
		}
		fmt.Printf("%v:%v %v\n", limitedCleaned, line, message)
	} else {
		fmt.Printf("unkown_file:? %v\n", message)
	}
}
