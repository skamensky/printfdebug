package runtime

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func Printf(message string, pathDepthFromEnd int) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		limited := filepath.Join(filepath.SplitList(file)[-pathDepthFromEnd:]...)
		fmt.Printf("%v:%v %v\n", limited, line, message)
	} else {
		fmt.Printf("unkown_file:? %v\n", message)

	}
}
