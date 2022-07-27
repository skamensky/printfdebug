package internal

import (
	"github.com/skamensky/printfdebug/internal/tests"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// all paths are from the perspective of the root project directory
	err := os.Chdir("..")
	if err != nil {
		panic(err)
	}
	code := m.Run()
	os.Exit(code)
}

func TestAnonymousFuncNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/AnonymousFunc.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestAnonymousFuncYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/AnonymousFunc.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/Func.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/Func.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncImmediatelyInvokedNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncImmediatelyInvoked.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncImmediatelyInvokedYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncImmediatelyInvoked.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncLogFatalNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncLogFatal.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncLogFatalYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncLogFatal.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncOsExitNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncOsExit.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncOsExitYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncOsExit.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncPanicNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncPanic.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncPanicYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncPanic.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncStoredAsVarNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncStoredAsVar.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncStoredAsVarYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncStoredAsVar.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncStoredAsVarAssignmentNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncStoredAsVarAssignment.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncStoredAsVarAssignmentYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncStoredAsVarAssignment.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncStoredAsVarValueSpecNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncStoredAsVarValueSpec.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncStoredAsVarValueSpecYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncStoredAsVarValueSpec.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncWithArgsAndReturnNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncWithArgsAndReturn.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncWithArgsAndReturnYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncWithArgsAndReturn.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncWithExplicitReturnNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncWithExplicitReturn.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncWithExplicitReturnYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncWithExplicitReturn.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncWithManyReturnStatementsNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncWithManyReturnStatements.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncWithManyReturnStatementsYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncWithManyReturnStatements.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncWithNestedAnonymousFuncsNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncWithNestedAnonymousFuncs.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestFuncWithNestedAnonymousFuncsYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/FuncWithNestedAnonymousFuncs.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestStructMethodNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/StructMethod.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestStructMethodYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/StructMethod.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestStructMethodWithArgsNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/StructMethodWithArgs.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestStructMethodWithArgsYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/StructMethodWithArgs.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestStructMethodWithArgsAndReturnNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/StructMethodWithArgsAndReturn.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestStructMethodWithArgsAndReturnYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/StructMethodWithArgsAndReturn.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestStructMethodWithReturnNoRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/StructMethodWithReturn.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}

func TestStructMethodWithReturnYesRuntime(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "internal/tests/testassets/StructMethodWithReturn.go",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}
