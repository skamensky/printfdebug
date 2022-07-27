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
