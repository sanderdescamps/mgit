package cmd

import (
	"github.com/spf13/cobra"
)

var cmdTest = &cobra.Command{
	Use:   "test",
	Short: "this is a test command",
	Long:  `This command doesn't do anything other than testing`,
	Run: func(cmd *cobra.Command, args []string) {
		display.Error("this is a ERROR message")
		display.Debug("this is a DEBUG message")
	},
}
