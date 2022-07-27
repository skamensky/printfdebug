package main

import (
	"fmt"
	"github.com/skamensky/printfdebug/internal/tests"
	"os"
	"path"
	"strings"
	"sync"
)

func runBuild() {
	// TODO
}

func genTestResults() {
	// TODO only append to path, don't overwrite
	os.Setenv("PATH", ".")
	wg := sync.WaitGroup{}
	errorChan := make(chan error)
	files := tests.GetTestFiles()
	for _, file := range files {
		wg.Add(1)
		go func(file tests.TestFile, errorChan chan error) {
			defer wg.Done()

			stderr, stdout, err := file.ExecuteCommand(file.BuildCommand(true))
			errorChan <- err

			if stderr != "" {
				fmt.Printf("stderr: %v\n", stdout)
			}
			if stdout != "" {
				fmt.Printf("stdout: %v\n", stdout)
			}
		}(file, errorChan)
	}
	errOccurred := false
	for i := 0; i < len(files); i++ {
		err := <-errorChan
		if err != nil {
			errOccurred = true
			fmt.Printf("Error: %v\n", err)
		}
	}
	wg.Wait()
	if !errOccurred {
		fmt.Println("Generated all test results with no failures")
	}
}

func writeTestFile() {
	fileHeader := `package internal

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
`
	files := tests.GetTestFiles()
	testsFunctions := []string{fileHeader}
	for _, file := range files {
		testsFunctions = append(testsFunctions, file.GenerateTestFunction())
	}
	functionDefs := []byte(strings.Join(testsFunctions, "\n"))

	testFilePath := path.Join("internal", "ast_test.go")

	err := os.WriteFile(testFilePath, functionDefs, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Generated all test cases")

}

func main() {
	runBuild()
	//genTestResults()
	writeTestFile()
}
