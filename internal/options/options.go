package options

import (
	"github.com/spf13/cobra"
)

type Options struct {
	Inplace  bool
	FilePath string
}

func NewOptionsFromCmd(cmd *cobra.Command) *Options {
	return NewOptions(
		true,
		cmd.Flags().Lookup("file").Value.String(),
	)

}

func NewOptions(inplace bool, filepath string) *Options {
	options := &Options{
		Inplace:  inplace,
		FilePath: "",
	}
	return options
}
