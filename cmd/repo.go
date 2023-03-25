package cmd

import (
	"github.com/spf13/cobra"
)

var cmdRepo = &cobra.Command{
	Use:   "repo",
	Short: "show repo config",
	Long:  `Manage repo configuration (default: show)`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdRepoShow.Run(cmd, args)
	},
	PreRun:  cmdRepoShow.PreRun,
	PostRun: cmdRepoShow.PostRun,
}

var cmdRepoShow = &cobra.Command{
	Use:   "show",
	Short: "show repo config",
	Long:  `Manage repo configuration (default: show)`,
	PreRun: func(cmd *cobra.Command, args []string) {
		initRepo()
	},
	Run: func(cmd *cobra.Command, args []string) {
		for idx, i := range repoCfg.GetRepoList() {
			display.Print(i.GitUrl)
			display.Infof("%d", idx)
		}
	},
}
