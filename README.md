# What is this?
A go tool to help with adding and removing `printf` debug statements

# Why?
I created this to learn golang and because while working on another project I was very confused by a wall of output produced by a `panic` from within a goroutine from within a library dependency.

I don't know how to use go debug tools.

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

# Examples
The below examples are generated automatically and are also the tests

# Prior work

https://github.com/yuuki0xff/printfdebug 

I actually only found this project after I wrote this tool and was looking for a name. yuuki0xff's printfdebug is not maintained and less complete.

yuuki0xff, If you're reading this let me know if you'd like to combine forces!