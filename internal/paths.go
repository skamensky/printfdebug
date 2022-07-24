package internal

import (
	"bytes"
	"github.com/skamensky/printfdebug/internal/options"
	"os"
)

func NormalizePath(filepath string) string {
	// TODO
	return filepath
}

func PathExists() {

}

func WriteOutput(options *options.Options, byteBuff *bytes.Buffer) error {
	if options.Inplace {
		return os.WriteFile(options.FilePath, byteBuff.Bytes(), 0644)
	} else {
		_, err := os.Stdout.Write(byteBuff.Bytes())
		return err
	}
}
