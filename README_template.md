# What is this?
A go tool to help with adding and removing `printf` debug statements

# Why?
I created this to learn golang and because while working on another project I was very confused by a wall of output produced by a `panic` from within a goroutine from within a library dependency.

I don't know how to use go debug tools.

# Usage

{INJECT_USAGE}

# Features
- Modifies your code by adding `printf` statements upon entering and exiting functions
- Captures all function exits branches, panics, system exits, and fatal logs
- Optionally injects a function to help print the current file name and line number (if anyone knows a way of doing this on a single line, I could remove this feature completely!)
- Support all functions (that I know of)
  - functions stored as variables (uses variable name as function name)
  - struct methods
  - anonymous functions
  - declared functions
  - nested functions
  - immediately invoked functions
- Idempotent operations.
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

{INJECT_EXAMPLES}

# Prior work

https://github.com/yuuki0xff/printfdebug 

I actually only found this project after I wrote this tool and was looking for a name. yuuki0xff's printfdebug is not maintained and less complete.

yuuki0xff, If you're reading this let me know if you'd like to combine forces!