/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"printfdebug/internal/options"

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
	rootCmd.PersistentFlags().BoolVarP(&OPTIONS.Inplace, "inplace", "i", false, "If specified, the file on disk is overwritten. The default is to output to stdout.")
	rootCmd.PersistentFlags().StringVarP(&OPTIONS.FilePath, "file", "f", "", "The file path.")
	_ = rootCmd.MarkPersistentFlagRequired("file")

}
