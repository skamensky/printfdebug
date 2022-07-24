package internal

import (
	"bytes"
	"github.com/skamensky/printfdebug/internal/options"
	"os"
	"path/filepath"
)

func NormalizePath(_filepatb string) string {
	// TODO handle error
	_path, _ := filepath.Abs(_filepatb)
	return _path
}

func PathExists() {
	// TODO
}

func WriteOutput(options *options.Options, byteBuff *bytes.Buffer) error {
	if options.Overwrite {
		return os.WriteFile(options.FilePath, byteBuff.Bytes(), 0644)
	} else if options.OutFile != "" {
		// TODO put this into command initialization. Handle for input file as well
		p := NormalizePath(options.OutFile)
		file, err := os.Create(p)
		defer file.Close()
		if err != nil {
			return err
		}
		_, err = file.Write(byteBuff.Bytes())
		if err != nil {
			return err
		}
		return nil
	} else {
		_, err := os.Stdout.Write(byteBuff.Bytes())
		return err
	}
}
