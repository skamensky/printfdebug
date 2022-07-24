package cmd

import (
	"bytes"
	"github.com/skamensky/printfdebug/internal"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add printf debug statements to all functions.",
	Long: `Add printf debug statements to all functions.
This is idempotent. If you run "add" on a file multiple times, the output will only change the first time.
When you run "add" on a file, it first removes all printf debug statements, and then adds again according to the given configuration.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.Open(OPTIONS.FilePath)

		defer file.Close()
		if err != nil {
			return err
		}
		byteBuffer := &bytes.Buffer{}
		_, err = io.Copy(byteBuffer, file)
		if err != nil {
			return err
		}
		removedBytesBuf, err := internal.RemovePrintDebugging(OPTIONS, byteBuffer)
		if err != nil {
			return err
		}
		byteBuf, err := internal.AddPrintDebugging(OPTIONS, removedBytesBuf)

		if err != nil {
			return err
		} else {
			return internal.WriteOutput(OPTIONS, byteBuf)
		}
	},
}
