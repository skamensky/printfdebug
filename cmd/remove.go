package cmd

import (
	"bytes"
	"github.com/spf13/cobra"
	"io"
	"os"
	"printfdebug/internal"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove printf debug statements in all locations that this tool previously added them to.",
	Long: `Remove printf debug statements in all locations that this tool previously added them to.
This relies on the identifier inserted as a commend after each function. It's therefore critical that you don't remove the commend or modify which line it's on. 
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
		return internal.WriteOutput(OPTIONS, removedBytesBuf)
	},
}
