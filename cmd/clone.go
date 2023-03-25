package cmd

import "github.com/spf13/cobra"

var cmdClone = &cobra.Command{
	Use:   "clone",
	Short: "clone repos",
	Long:  `Create the local directory structure and clone all repositories`,
	PreRun: func(cmd *cobra.Command, args []string) {
		initRepo()
	},
	Run: func(cmd *cobra.Command, args []string) {
		display.Info("clone repos")
	},
}
