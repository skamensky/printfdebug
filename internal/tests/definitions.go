package tests

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

type Command int
type TestType int

const (
	ADD Command = iota
	REMOVE
)
const (
	NoRuntime TestType = iota
	YesRuntime
)

type TestFile struct {
	Command Command
	InFile  string
	Type    TestType
}

func (t *TestFile) GetOutputFilePath() string {
	if t.Type == NoRuntime {
		return path.Join("internal", "tests", "testassets", "results", "noruntime", path.Base(t.InFile))
	} else {
		return path.Join("internal", "tests", "testassets", "results", "runtime", path.Base(t.InFile))
	}
}
func (t *TestFile) GetOutputFileContents() string {
	var contents []byte
	var err error
	contents, err = os.ReadFile(t.GetOutputFilePath())
	if err != nil {
		panic(err)
	}
	return string(contents)
}
func (t *TestFile) GetInputFileContents() string {
	contents, err := os.ReadFile(t.InFile)
	if err != nil {
		panic(err)
	}
	return string(contents)
}

func (t *TestFile) ExecuteCommand(commands []string) (string, string, error) {
	stderr := ""
	var err error = nil
	executable := "printfdebug"
	if strings.ToLower(os.Getenv("DEBUG")) == "true" {
		fmt.Printf("Executing command: \"%v %v\"\n", executable, strings.Join(commands, " "))

	}
	cmd := exec.Command(executable, commands...)
	stdout, err := cmd.Output()

	ee, ok := err.(*exec.ExitError)

	if ok {
		stderr = string(ee.Stderr)
		err = nil
	}
	return string(stdout), stderr, err
}

func (t *TestFile) BuildCommand(write bool) []string {
	commands := []string{"--file", t.InFile}
	if t.Command == ADD {
		commands = append(commands, "add")

	} else {
		commands = append(commands, "remove")
	}

	if write {
		commands = append(commands, []string{"--out-file", t.GetOutputFilePath()}...)
	}
	if t.Type == NoRuntime {
		commands = append(commands, "--no-runtime")
	}
	return commands
}

func (t *TestFile) GetFuncName() string {
	return strings.ReplaceAll(path.Base(t.InFile), ".go", "")
}

func (t *TestFile) GenerateTestFunction() string {
	typeComponent := "YesRuntime"
	if t.Type == NoRuntime {
		typeComponent = "NoRuntime"
	}
	// TODO path so it doesn't override but rather appends
	funcName := t.GetFuncName() + typeComponent
	return fmt.Sprintf(`func Test%v(t *testing.T) {
	
	err := os.Setenv("PATH",".")
	if err != nil {
		panic(err)
	}
	testFile := tests.TestFile{
		Command: tests.ADD,
		InFile:  "%v",
		Type:    tests.NoRuntime,
	}

	testFileRemove := testFile
	testFileRemove.Command = tests.REMOVE

	stdout, stderr, err := testFile.ExecuteCommand(testFile.BuildCommand(false))
	if stderr != "" {
		t.Errorf("Add command stderr is not empty: %%v", stderr)
	}
	if err != nil {
		t.Errorf("Add command resulted in an error %%v", err)
	}
	if testFile.GetOutputFileContents() != stdout {
		t.Errorf("%%v did not evaluate to the expected output result upon the add command", testFile.InFile)
	}

	stdout, stderr, err = testFileRemove.ExecuteCommand(testFileRemove.BuildCommand(false))

	if stderr != "" {
		t.Errorf("Remove command stderr is not empty: %%v", stderr)
	}
	if err != nil {
		t.Errorf("Remove command resulted in an error %%v", err)
	}
	if stdout != testFileRemove.GetInputFileContents() {
		t.Errorf("%%v did not reset to its original input upon the remove command", testFileRemove.InFile)
	}
}
`, funcName, t.InFile)
}

func GetTestFiles() []TestFile {
	testAssetDir := path.Join("internal", "tests", "testassets")
	entries, err := os.ReadDir(testAssetDir)
	testFiles := []TestFile{}
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") && entry.Name() != "definitions.go" {
			baseTest := TestFile{
				InFile:  path.Join(testAssetDir, entry.Name()),
				Command: ADD,
			}

			noRuntime := baseTest
			noRuntime.Type = NoRuntime

			yesRuntime := baseTest
			yesRuntime.Type = YesRuntime

			testFiles = append(testFiles, noRuntime)
			testFiles = append(testFiles, yesRuntime)
		}
	}
	return testFiles
}
