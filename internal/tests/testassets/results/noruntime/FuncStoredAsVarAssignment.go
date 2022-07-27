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
