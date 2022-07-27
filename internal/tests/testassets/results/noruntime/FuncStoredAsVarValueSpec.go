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
