package testassest

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func FuncStoredAsVarAssignment() {
	printfdebug_Printf_FuncStoredAsVarAssignment("Entering \"FuncStoredAsVarAssignment\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	FuncStoredAsVarInnerSingle := func() error {
		printfdebug_Printf_FuncStoredAsVarAssignment("Entering \"FuncStoredAsVarInnerSingle\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		printfdebug_Printf_FuncStoredAsVarAssignment("Leaving \"FuncStoredAsVarInnerSingle\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		return nil
	}
	FuncStoredAsVarInnerMulti, _ := func() error {
		printfdebug_Printf_FuncStoredAsVarAssignment("Entering \"FuncStoredAsVarInnerMulti\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		printfdebug_Printf_FuncStoredAsVarAssignment("Leaving \"FuncStoredAsVarInnerMulti\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		return nil
	}, ""

	_ = FuncStoredAsVarInnerMulti
	_ = FuncStoredAsVarInnerSingle
	printfdebug_Printf_FuncStoredAsVarAssignment("Leaving \"FuncStoredAsVarAssignment\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

var _ = runtime.Caller   // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = filepath.Clean   // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = fmt.Println      // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = strings.Split    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = os.PathSeparator // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_FuncStoredAsVarAssignment(message string, pathDepthFromEnd int) {
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
