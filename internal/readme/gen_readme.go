package main

import (
	"fmt"
	"github.com/skamensky/printfdebug/internal/tests"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var backTick = "`"
var tripBackTick = "```"

func SplitByCamcel(src string) []string {
	// from https://www.golangprograms.com/split-a-string-at-uppercase-letters-using-regular-expression-in-golang.html
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	results := []string{}

	submatchall := re.FindAllString(src, -1)
	for _, element := range submatchall {
		results = append(results, element)
	}
	return results
}

func GetMarkdownBlockForFunction(file tests.TestFile) string {
	cmd := file.BuildCommand(false)
	inContent := file.GetInputFileContents()
	outContent := file.GetOutputFileContents()
	words := SplitByCamcel(file.GetFuncName())
	if file.Type == tests.NoRuntime {
		words = append(words, "(No Runtime)")
	} else {
		words = append(words, "(With Runtime)")
	}
	for idx, word := range words {
		if word == "Func" {
			words[idx] = "Function"
		}
		if word == "Funcs" {
			words[idx] = "Functions"
		}
		if word == "Args" {
			words[idx] = "Arguments§"
		}
	}
	exampleName := strings.Join(words, " ")
	cmdNice := fmt.Sprintf("printfdebug %v", strings.Join(cmd, " "))

	return fmt.Sprintf(`<details>
  <summary>%v</summary>

Running %v%v%v
On a file containing
%vgo
%v
%v
Will produce the following result
%vgo
%v
%v
</details>`, exampleName, backTick, cmdNice, backTick, tripBackTick, inContent, tripBackTick, tripBackTick, outContent, tripBackTick)
}

func GetMarkdownForCLIUsage() string {
	cmd := exec.Command("./printfdebug")
	stderr := ""
	stdout, err := cmd.Output()

	ee, ok := err.(*exec.ExitError)

	if ok {
		stderr = string(ee.Stderr)
		err = nil
	}
	if stderr != "" {
		panic(stderr)
	} else if err != nil {
		panic(err)
	}
	return fmt.Sprintf(`CLI Usage
%vcmd
%v
%v`, tripBackTick, string(stdout), tripBackTick)
}

func main() {
	files := tests.GetTestFiles()
	examplesMarkdown := []string{}
	for _, f := range files {
		examplesMarkdown = append(examplesMarkdown, GetMarkdownBlockForFunction(f))
	}
	cliUsageMdown := GetMarkdownForCLIUsage()

	template, err := os.ReadFile("README_template.md")
	if err != nil {
		panic(err)
	}
	newMarkdown := string(template)
	newMarkdown = strings.ReplaceAll(newMarkdown, "{INJECT_EXAMPLES}", strings.Join(examplesMarkdown, "\n"))
	newMarkdown = strings.ReplaceAll(newMarkdown, "{INJECT_USAGE}", cliUsageMdown)

	err = os.WriteFile("README.md", []byte(newMarkdown), 0644)
	if err != nil {
		panic(err)
	}

}
