/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/skamensky/printfdebug/internal/options"
	"os"

	"github.com/spf13/cobra"
)

var OPTIONS = &options.Options{}

var rootCmd = &cobra.Command{
	Use:   "printfdebug",
	Short: "A brief description of your application",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&OPTIONS.Overwrite, "write", "w", false, "If specified, the file on disk is overwritten. The default is to output to stdout.")
	rootCmd.PersistentFlags().BoolVarP(&OPTIONS.NoRuntime, "no-runtime", "n", false, "Disable the injection of the printfdebug function definition into your file. The alternative is a simple fmt.Printf statement.")
	rootCmd.PersistentFlags().StringVarP(&OPTIONS.FilePath, "file", "f", "", "The file path.")
	rootCmd.PersistentFlags().StringVarP(&OPTIONS.OutFile, "out-file", "o", "", "An optional output file. Cannot be used with `write`")
	rootCmd.PersistentFlags().IntVarP(&OPTIONS.PathDepth, "path-depth", "d", 1, "Is only taken into account if `no-runtime` is not set. The depth of the directory tree to print from the printf function. Defaults to only the current file name.")
	_ = rootCmd.MarkPersistentFlagRequired("file")
	rootCmd.MarkFlagsMutuallyExclusive("out-file", "write")

}
