# What is this?
A go tool to help with adding and removing `printf` debug statements

# Why?
I created this to learn golang and because while working on another project I was very confused by a wall of output produced by a `panic` from within a goroutine from within a library dependency.

I don't know how to use go debug tools.

# Usage

CLI Usage
```cmd
Organized debugging using printf statements

Usage:
  printfdebug [command]

Available Commands:
  add         Add printf debug statements to all functions.
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  remove      Remove printf debug statements in all locations that this tool previously added them to.

Flags:
  -f, --file string             The file path.
  -h, --help                    help for printfdebug
  -n, --no-runtime              Disable the injection of the printfdebug function definition into your file. The alternative is a simple fmt.Printf statement.
  -o, --out-file write          An optional output file. Cannot be used with write
  -d, --path-depth no-runtime   Is only taken into account if no-runtime is not set. The depth of the directory tree to print from the printf function. Defaults to only the current file name. (default 1)
  -w, --write                   If specified, the file on disk is overwritten. The default is to output to stdout.

Use "printfdebug [command] --help" for more information about a command.

```

# Features
- Modifies your code by adding `printf` statements upon entering and exiting functions
- Captures all function exits branches, panics, system exits, and fatal logs
- Optionally injects a function to help print the current file name and line number (if anyone knows a way of doing this on a single line, I could remove this feature completely!)
- Support all functions (that I know of)
  - functions stored as variables
  - struct methods
  - anonymous functions
  - declared functions
  - nested functions
  Idempotent operations.
  - Running `printfdebug add` multiple times on the same input produces sane results
  - Running `printfdebug add` on code that's been modified since the last run updates the code to add new `printf` statements
  - Running `printfdebug remove` cleans up all traces of the debugger

# Non features
- Complex control flow analysis to avoid extraneous placement of printf statements
- Working with custom code bases that have modified the standard library

# Todo
- Go fmt automatically (or does dst alreadu handle that?)
- Optional message formatting (before and after function with linenum and name as template vars)
- Optional list of excluded or included function names
- Clean up TODO's in code

# Testing
To add a new test

1. Make your changes
2. Create a new file in [internal/testassets](internal/tests/testassets)
3. Run `go run internal/tests/gen_tests/gen_tests.go` to generate the output according to the current state of the package. This will:
   1. Do a rebuild of the cli and overwrite the file `printfdebug` in the root of the repo
   2. Generate output for each file in  [internal/tests/testassets](internal/tests/testassets) for the standard case and `no-runtime` case
   3. Generate a test for each case and write to  [internal/ast_test.go](internal/ast_test.go)
4. Make sure the outputs in  [internal/tests/testassets/results/noruntime](internal/tests/testassets/results/noruntime) and  [internal/tests/testassets/results/runtime](internal/tests/testassets/results/runtime) are correct (otherwise go back to step 1 and iterate)
5. Run `go test -v internal/ast_test.go` from the root directory (right now things are very path dependent)

# Readme

The readme is generated automatically to inject examples and CLI usage.
After adding tests or modifying the CLI, run ` go run internal/readme/gen_readme.go`

Do not make changed directly to the `README.md` file. Only ever make changes to the file `README_template.md` and regenerate.
# Examples
The below examples are generated automatically from test results

<details>
  <summary>Anonymous Function (No Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/AnonymousFunc.go add --no-runtime`
On a file containing
```go
package testassest

func AnonymousFunc() {
	func() {
	}()
}

```
Will produce the following result
```go
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

```
</details>
<details>
  <summary>Anonymous Function (With Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/AnonymousFunc.go add`
On a file containing
```go
package testassest

func AnonymousFunc() {
	func() {
	}()
}

```
Will produce the following result
```go
package testassest

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func AnonymousFunc() {
	printfdebug_Printf_AnonymousFunc("Entering \"AnonymousFunc\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	func() {
		printfdebug_Printf_AnonymousFunc("Entering \"anonymous-function Fun\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		printfdebug_Printf_AnonymousFunc("Leaving \"anonymous-function Fun\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	}()
	printfdebug_Printf_AnonymousFunc("Leaving \"AnonymousFunc\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

var _ = runtime.Caller // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = filepath.Clean // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = fmt.Println    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_AnonymousFunc(message string, pathDepthFromEnd int) {
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

```
</details>
<details>
  <summary>Function (No Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/Func.go add --no-runtime`
On a file containing
```go
package testassest

func Func() {
}

```
Will produce the following result
```go
package testassest

import "fmt"

func Func() {
	fmt.Println("Entering \"Func\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"Func\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

```
</details>
<details>
  <summary>Function (With Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/Func.go add`
On a file containing
```go
package testassest

func Func() {
}

```
Will produce the following result
```go
package testassest

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func Func() {
	printfdebug_Printf_Func("Entering \"Func\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	printfdebug_Printf_Func("Leaving \"Func\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

var _ = runtime.Caller // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = filepath.Clean // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = fmt.Println    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_Func(message string, pathDepthFromEnd int) {
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

```
</details>
<details>
  <summary>Function Log Fatal (No Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncLogFatal.go add --no-runtime`
On a file containing
```go
package testassest

import (
	"log"
)

func FuncLogFatal() {
	log.Fatal(1)
}

func FuncLogFatalln() {
	log.Fatalln(1)

}
func FuncLogFatalf() {
	log.Fatalf("")

}

```
Will produce the following result
```go
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

```
</details>
<details>
  <summary>Function Log Fatal (With Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncLogFatal.go add`
On a file containing
```go
package testassest

import (
	"log"
)

func FuncLogFatal() {
	log.Fatal(1)
}

func FuncLogFatalln() {
	log.Fatalln(1)

}
func FuncLogFatalf() {
	log.Fatalf("")

}

```
Will produce the following result
```go
package testassest

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
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

var _ = runtime.Caller // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = filepath.Clean // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = fmt.Println    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_FuncLogFatal(message string, pathDepthFromEnd int) {
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

```
</details>
<details>
  <summary>Function Os Exit (No Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncOsExit.go add --no-runtime`
On a file containing
```go
package testassest

import "os"

func FuncOsExit() {
	os.Exit(1)
}

```
Will produce the following result
```go
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

```
</details>
<details>
  <summary>Function Os Exit (With Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncOsExit.go add`
On a file containing
```go
package testassest

import "os"

func FuncOsExit() {
	os.Exit(1)
}

```
Will produce the following result
```go
package testassest

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func FuncOsExit() {
	printfdebug_Printf_FuncOsExit("Entering \"FuncOsExit\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	printfdebug_Printf_FuncOsExit("Leaving \"FuncOsExit\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	os.Exit(1)
	printfdebug_Printf_FuncOsExit("Leaving \"FuncOsExit\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

var _ = filepath.Clean // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = fmt.Println    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = runtime.Caller // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_FuncOsExit(message string, pathDepthFromEnd int) {
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

```
</details>
<details>
  <summary>Function Panic (No Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncPanic.go add --no-runtime`
On a file containing
```go
package testassest

func FuncPanic() {
	panic("Oh no, how will we know how we got here?")
	return
}

```
Will produce the following result
```go
package testassest

import "fmt"

func FuncPanic() {
	fmt.Println("Entering \"FuncPanic\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"FuncPanic\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	panic("Oh no, how will we know how we got here?")
	fmt.Println("Leaving \"FuncPanic\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	return
}

```
</details>
<details>
  <summary>Function Panic (With Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncPanic.go add`
On a file containing
```go
package testassest

func FuncPanic() {
	panic("Oh no, how will we know how we got here?")
	return
}

```
Will produce the following result
```go
package testassest

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func FuncPanic() {
	printfdebug_Printf_FuncPanic("Entering \"FuncPanic\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	printfdebug_Printf_FuncPanic("Leaving \"FuncPanic\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	panic("Oh no, how will we know how we got here?")
	printfdebug_Printf_FuncPanic("Leaving \"FuncPanic\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	return
}

var _ = runtime.Caller // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = filepath.Clean // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = fmt.Println    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_FuncPanic(message string, pathDepthFromEnd int) {
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

```
</details>
<details>
  <summary>Function Stored As Var (No Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncStoredAsVar.go add --no-runtime`
On a file containing
```go
package testassest

var FuncStoredAsVar = func() error {
	return nil
}

```
Will produce the following result
```go
package testassest

import "fmt"

var FuncStoredAsVar = func() error {
	fmt.Println("Entering \"FuncStoredAsVar\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"FuncStoredAsVar\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	return nil
}

```
</details>
<details>
  <summary>Function Stored As Var (With Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncStoredAsVar.go add`
On a file containing
```go
package testassest

var FuncStoredAsVar = func() error {
	return nil
}

```
Will produce the following result
```go
package testassest

import (
	"fmt"
	"path/filepath"
	"runtime"
)

var FuncStoredAsVar = func() error {
	printfdebug_Printf_FuncStoredAsVar("Entering \"FuncStoredAsVar\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	printfdebug_Printf_FuncStoredAsVar("Leaving \"FuncStoredAsVar\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	return nil
}
var _ = runtime.Caller // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = filepath.Clean // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = fmt.Println    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_FuncStoredAsVar(message string, pathDepthFromEnd int) {
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

```
</details>
<details>
  <summary>Function Stored As Var Assignment (No Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncStoredAsVarAssignment.go add --no-runtime`
On a file containing
```go
package testassest

func FuncStoredAsVarAssignment() {
	FuncStoredAsVarInnerSingle := func() error {
		return nil
	}
	FuncStoredAsVarInnerMulti, _ := func() error {
		return nil
	}, ""

	_ = FuncStoredAsVarInnerMulti
	_ = FuncStoredAsVarInnerSingle
}

```
Will produce the following result
```go
package testassest

import "fmt"

func FuncStoredAsVarAssignment() {
	fmt.Println("Entering \"FuncStoredAsVarAssignment\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	FuncStoredAsVarInnerSingle := func() error {
		fmt.Println("Entering \"FuncStoredAsVarInnerSingle\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		fmt.Println("Leaving \"FuncStoredAsVarInnerSingle\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		return nil
	}
	FuncStoredAsVarInnerMulti, _ := func() error {
		fmt.Println("Entering \"FuncStoredAsVarInnerMulti\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		fmt.Println("Leaving \"FuncStoredAsVarInnerMulti\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		return nil
	}, ""

	_ = FuncStoredAsVarInnerMulti
	_ = FuncStoredAsVarInnerSingle
	fmt.Println("Leaving \"FuncStoredAsVarAssignment\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

```
</details>
<details>
  <summary>Function Stored As Var Assignment (With Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncStoredAsVarAssignment.go add`
On a file containing
```go
package testassest

func FuncStoredAsVarAssignment() {
	FuncStoredAsVarInnerSingle := func() error {
		return nil
	}
	FuncStoredAsVarInnerMulti, _ := func() error {
		return nil
	}, ""

	_ = FuncStoredAsVarInnerMulti
	_ = FuncStoredAsVarInnerSingle
}

```
Will produce the following result
```go
package testassest

import (
	"fmt"
	"path/filepath"
	"runtime"
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

var _ = runtime.Caller // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = filepath.Clean // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = fmt.Println    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_FuncStoredAsVarAssignment(message string, pathDepthFromEnd int) {
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

```
</details>
<details>
  <summary>Function Stored As Var Value Spec (No Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncStoredAsVarValueSpec.go add --no-runtime`
On a file containing
```go
package testassest

func FuncStoredAsVarValueSpec() {
	var FuncStoredAsVarInnerSingle = func() error {
		return nil
	}

	var FuncStoredAsVarInnerMutli, _ = func() error {
		return nil
	}, ""

	_ = FuncStoredAsVarInnerSingle
	_ = FuncStoredAsVarInnerMutli
}

```
Will produce the following result
```go
package testassest

import "fmt"

func FuncStoredAsVarValueSpec() {
	fmt.Println("Entering \"FuncStoredAsVarValueSpec\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	var FuncStoredAsVarInnerSingle = func() error {
		fmt.Println("Entering \"FuncStoredAsVarInnerSingle\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		fmt.Println("Leaving \"FuncStoredAsVarInnerSingle\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		return nil
	}

	var FuncStoredAsVarInnerMutli, _ = func() error {
		fmt.Println("Entering \"FuncStoredAsVarInnerMutli\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		fmt.Println("Leaving \"FuncStoredAsVarInnerMutli\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		return nil
	}, ""

	_ = FuncStoredAsVarInnerSingle
	_ = FuncStoredAsVarInnerMutli
	fmt.Println("Leaving \"FuncStoredAsVarValueSpec\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

```
</details>
<details>
  <summary>Function Stored As Var Value Spec (With Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncStoredAsVarValueSpec.go add`
On a file containing
```go
package testassest

func FuncStoredAsVarValueSpec() {
	var FuncStoredAsVarInnerSingle = func() error {
		return nil
	}

	var FuncStoredAsVarInnerMutli, _ = func() error {
		return nil
	}, ""

	_ = FuncStoredAsVarInnerSingle
	_ = FuncStoredAsVarInnerMutli
}

```
Will produce the following result
```go
package testassest

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func FuncStoredAsVarValueSpec() {
	printfdebug_Printf_FuncStoredAsVarValueSpec("Entering \"FuncStoredAsVarValueSpec\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	var FuncStoredAsVarInnerSingle = func() error {
		printfdebug_Printf_FuncStoredAsVarValueSpec("Entering \"FuncStoredAsVarInnerSingle\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		printfdebug_Printf_FuncStoredAsVarValueSpec("Leaving \"FuncStoredAsVarInnerSingle\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		return nil
	}

	var FuncStoredAsVarInnerMutli, _ = func() error {
		printfdebug_Printf_FuncStoredAsVarValueSpec("Entering \"FuncStoredAsVarInnerMutli\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		printfdebug_Printf_FuncStoredAsVarValueSpec("Leaving \"FuncStoredAsVarInnerMutli\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		return nil
	}, ""

	_ = FuncStoredAsVarInnerSingle
	_ = FuncStoredAsVarInnerMutli
	printfdebug_Printf_FuncStoredAsVarValueSpec("Leaving \"FuncStoredAsVarValueSpec\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

var _ = runtime.Caller // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = filepath.Clean // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = fmt.Println    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_FuncStoredAsVarValueSpec(message string, pathDepthFromEnd int) {
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

```
</details>
<details>
  <summary>Function With Arguments§ And Return (No Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncWithArgsAndReturn.go add --no-runtime`
On a file containing
```go
package testassest

func FuncWithArgsAndReturn(s2 string) error {
	// this is a comment on its own line
	/*
		This is a multi line
		comment
	*/
	return nil //This is inline comment
}

```
Will produce the following result
```go
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

```
</details>
<details>
  <summary>Function With Arguments§ And Return (With Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncWithArgsAndReturn.go add`
On a file containing
```go
package testassest

func FuncWithArgsAndReturn(s2 string) error {
	// this is a comment on its own line
	/*
		This is a multi line
		comment
	*/
	return nil //This is inline comment
}

```
Will produce the following result
```go
package testassest

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func FuncWithArgsAndReturn(s2 string) error {
	printfdebug_Printf_FuncWithArgsAndReturn("Entering \"FuncWithArgsAndReturn\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	printfdebug_Printf_FuncWithArgsAndReturn("Leaving \"FuncWithArgsAndReturn\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	// this is a comment on its own line
	/*
		This is a multi line
		comment
	*/
	return nil //This is inline comment
}

var _ = fmt.Println    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = runtime.Caller // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = filepath.Clean // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_FuncWithArgsAndReturn(message string, pathDepthFromEnd int) {
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

```
</details>
<details>
  <summary>Function With Explicit Return (No Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncWithExplicitReturn.go add --no-runtime`
On a file containing
```go
package testassest

func FuncWithExplicitReturn() {
	return
}

```
Will produce the following result
```go
package testassest

import "fmt"

func FuncWithExplicitReturn() {
	fmt.Println("Entering \"FuncWithExplicitReturn\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"FuncWithExplicitReturn\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	return
}

```
</details>
<details>
  <summary>Function With Explicit Return (With Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncWithExplicitReturn.go add`
On a file containing
```go
package testassest

func FuncWithExplicitReturn() {
	return
}

```
Will produce the following result
```go
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

```
</details>
<details>
  <summary>Function With Many Return Statements (No Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncWithManyReturnStatements.go add --no-runtime`
On a file containing
```go
package testassest

func FuncWithManyReturnStatements() {
	if false {
		return
	}
	if false {
		return
	}
	if false {
		return
	}
}

```
Will produce the following result
```go
package testassest

import "fmt"

func FuncWithManyReturnStatements() {
	fmt.Println("Entering \"FuncWithManyReturnStatements\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	if false {
		fmt.Println("Leaving \"FuncWithManyReturnStatements\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		return
	}
	if false {
		fmt.Println("Leaving \"FuncWithManyReturnStatements\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		return
	}
	if false {
		fmt.Println("Leaving \"FuncWithManyReturnStatements\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		return
	}
	fmt.Println("Leaving \"FuncWithManyReturnStatements\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

```
</details>
<details>
  <summary>Function With Many Return Statements (With Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncWithManyReturnStatements.go add`
On a file containing
```go
package testassest

func FuncWithManyReturnStatements() {
	if false {
		return
	}
	if false {
		return
	}
	if false {
		return
	}
}

```
Will produce the following result
```go
package testassest

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func FuncWithManyReturnStatements() {
	printfdebug_Printf_FuncWithManyReturnStatements("Entering \"FuncWithManyReturnStatements\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	if false {
		printfdebug_Printf_FuncWithManyReturnStatements("Leaving \"FuncWithManyReturnStatements\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		return
	}
	if false {
		printfdebug_Printf_FuncWithManyReturnStatements("Leaving \"FuncWithManyReturnStatements\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		return
	}
	if false {
		printfdebug_Printf_FuncWithManyReturnStatements("Leaving \"FuncWithManyReturnStatements\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		return
	}
	printfdebug_Printf_FuncWithManyReturnStatements("Leaving \"FuncWithManyReturnStatements\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

var _ = runtime.Caller // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = filepath.Clean // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = fmt.Println    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_FuncWithManyReturnStatements(message string, pathDepthFromEnd int) {
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

```
</details>
<details>
  <summary>Function With Nested Anonymous Functions (No Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncWithNestedAnonymousFuncs.go add --no-runtime`
On a file containing
```go
package testassest

func FuncWithNestedAnonymousFuncs() {
	innerFunc := func() {
		nestedInnerFunc := func() {}
		nestedInnerFunc()
	}
	_ = innerFunc
}

```
Will produce the following result
```go
package testassest

import "fmt"

func FuncWithNestedAnonymousFuncs() {
	fmt.Println("Entering \"FuncWithNestedAnonymousFuncs\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	innerFunc := func() {
		fmt.Println("Entering \"innerFunc\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		nestedInnerFunc := func() {
			fmt.Println("Entering \"nestedInnerFunc\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
			fmt.Println("Leaving \"nestedInnerFunc\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		}
		nestedInnerFunc()
		fmt.Println("Leaving \"innerFunc\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	}
	_ = innerFunc
	fmt.Println("Leaving \"FuncWithNestedAnonymousFuncs\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

```
</details>
<details>
  <summary>Function With Nested Anonymous Functions (With Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/FuncWithNestedAnonymousFuncs.go add`
On a file containing
```go
package testassest

func FuncWithNestedAnonymousFuncs() {
	innerFunc := func() {
		nestedInnerFunc := func() {}
		nestedInnerFunc()
	}
	_ = innerFunc
}

```
Will produce the following result
```go
package testassest

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func FuncWithNestedAnonymousFuncs() {
	printfdebug_Printf_FuncWithNestedAnonymousFuncs("Entering \"FuncWithNestedAnonymousFuncs\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	innerFunc := func() {
		printfdebug_Printf_FuncWithNestedAnonymousFuncs("Entering \"innerFunc\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		nestedInnerFunc := func() {
			printfdebug_Printf_FuncWithNestedAnonymousFuncs("Entering \"nestedInnerFunc\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
			printfdebug_Printf_FuncWithNestedAnonymousFuncs("Leaving \"nestedInnerFunc\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
		}
		nestedInnerFunc()
		printfdebug_Printf_FuncWithNestedAnonymousFuncs("Leaving \"innerFunc\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	}
	_ = innerFunc
	printfdebug_Printf_FuncWithNestedAnonymousFuncs("Leaving \"FuncWithNestedAnonymousFuncs\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

var _ = runtime.Caller // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = filepath.Clean // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = fmt.Println    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_FuncWithNestedAnonymousFuncs(message string, pathDepthFromEnd int) {
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

```
</details>
<details>
  <summary>Struct Method (No Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/StructMethod.go add --no-runtime`
On a file containing
```go
package testassest

func (s *S) StructMethod() {

}

```
Will produce the following result
```go
package testassest

import "fmt"

func (s *S) StructMethod() {

	fmt.Println("Entering \"StructMethod\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"StructMethod\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

```
</details>
<details>
  <summary>Struct Method (With Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/StructMethod.go add`
On a file containing
```go
package testassest

func (s *S) StructMethod() {

}

```
Will produce the following result
```go
package testassest

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func (s *S) StructMethod() {

	printfdebug_Printf_StructMethod("Entering \"StructMethod\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	printfdebug_Printf_StructMethod("Leaving \"StructMethod\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

var _ = filepath.Clean // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = fmt.Println    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = runtime.Caller // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_StructMethod(message string, pathDepthFromEnd int) {
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

```
</details>
<details>
  <summary>Struct Method With Arguments§ (No Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/StructMethodWithArgs.go add --no-runtime`
On a file containing
```go
package testassest

func (s *S) StructMethodWithArgs(s2 string) {
}

```
Will produce the following result
```go
package testassest

import "fmt"

func (s *S) StructMethodWithArgs(s2 string) {
	fmt.Println("Entering \"StructMethodWithArgs\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"StructMethodWithArgs\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

```
</details>
<details>
  <summary>Struct Method With Arguments§ (With Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/StructMethodWithArgs.go add`
On a file containing
```go
package testassest

func (s *S) StructMethodWithArgs(s2 string) {
}

```
Will produce the following result
```go
package testassest

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func (s *S) StructMethodWithArgs(s2 string) {
	printfdebug_Printf_StructMethodWithArgs("Entering \"StructMethodWithArgs\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	printfdebug_Printf_StructMethodWithArgs("Leaving \"StructMethodWithArgs\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
}

var _ = runtime.Caller // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = filepath.Clean // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = fmt.Println    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_StructMethodWithArgs(message string, pathDepthFromEnd int) {
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

```
</details>
<details>
  <summary>Struct Method With Arguments§ And Return (No Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/StructMethodWithArgsAndReturn.go add --no-runtime`
On a file containing
```go
package testassest

func (s *S) StructMethodWithArgsAndReturn(s2 string) error {
	return nil
}

```
Will produce the following result
```go
package testassest

import "fmt"

func (s *S) StructMethodWithArgsAndReturn(s2 string) error {
	fmt.Println("Entering \"StructMethodWithArgsAndReturn\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"StructMethodWithArgsAndReturn\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	return nil
}

```
</details>
<details>
  <summary>Struct Method With Arguments§ And Return (With Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/StructMethodWithArgsAndReturn.go add`
On a file containing
```go
package testassest

func (s *S) StructMethodWithArgsAndReturn(s2 string) error {
	return nil
}

```
Will produce the following result
```go
package testassest

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func (s *S) StructMethodWithArgsAndReturn(s2 string) error {
	printfdebug_Printf_StructMethodWithArgsAndReturn("Entering \"StructMethodWithArgsAndReturn\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	printfdebug_Printf_StructMethodWithArgsAndReturn("Leaving \"StructMethodWithArgsAndReturn\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	return nil
}

var _ = fmt.Println    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = runtime.Caller // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = filepath.Clean // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_StructMethodWithArgsAndReturn(message string, pathDepthFromEnd int) {
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

```
</details>
<details>
  <summary>Struct Method With Return (No Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/StructMethodWithReturn.go add --no-runtime`
On a file containing
```go
package testassest

func (s *S) StructMethodWithReturn() error {
	return nil
}

```
Will produce the following result
```go
package testassest

import "fmt"

func (s *S) StructMethodWithReturn() error {
	fmt.Println("Entering \"StructMethodWithReturn\"\n") // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	fmt.Println("Leaving \"StructMethodWithReturn\"\n")  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	return nil
}

```
</details>
<details>
  <summary>Struct Method With Return (With Runtime)</summary>

Running `printfdebug --file internal/tests/testassets/StructMethodWithReturn.go add`
On a file containing
```go
package testassest

func (s *S) StructMethodWithReturn() error {
	return nil
}

```
Will produce the following result
```go
package testassest

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func (s *S) StructMethodWithReturn() error {
	printfdebug_Printf_StructMethodWithReturn("Entering \"StructMethodWithReturn\"\n", 1) // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	printfdebug_Printf_StructMethodWithReturn("Leaving \"StructMethodWithReturn\"\n", 1)  // automatically added by printf-debugger. Do not change this comment. It is an identifier.
	return nil
}

var _ = runtime.Caller // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = filepath.Clean // automatically added by printf-debugger. Do not change this comment. It is an identifier.
var _ = fmt.Println    // automatically added by printf-debugger. Do not change this comment. It is an identifier.
func printfdebug_Printf_StructMethodWithReturn(message string, pathDepthFromEnd int) {
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

```
</details>

# Prior work

https://github.com/yuuki0xff/printfdebug 

I actually only found this project after I wrote this tool and was looking for a name. yuuki0xff's printfdebug is not maintained and less complete.

yuuki0xff, If you're reading this let me know if you'd like to combine forces!