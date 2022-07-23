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
