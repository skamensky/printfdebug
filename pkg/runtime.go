package runtime

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func maxInt(first int, second int) (max int) {
	if first > second {
		return first
	} else {
		return second
	}
}

func Printf(message string, pathDepthFromEnd int) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fileParts := filepath.SplitList(file)
		pathFromEndSafe := maxInt(len(fileParts), pathDepthFromEnd)
		limited := filepath.Join(fileParts[pathFromEndSafe:]...)
		fmt.Printf("%v:%v %v\n", limited, line, message)
	} else {
		fmt.Printf("unkown_file:? %v\n", message)

	}
}
