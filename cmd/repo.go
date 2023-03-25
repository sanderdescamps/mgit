package cmd

import (
	"strings"

	"github.com/sanderdescamps/mgit/internal/config"
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
		display.Print(strings.Repeat("-", 80))
		for _, i := range repoCfg.GetRepoList() {
			display.Print("%s", i.GetRepoCloneUrl())
			display.Print(" path: %s", i.GetFSPath())
			display.Print(" settings:")
			for key, val := range i.GetAllSettings() {
				display.Print("   - %s: %v", key, val)
			}
			display.Print(strings.Repeat("-", 80))
		}
	},
}

		}
	},
}
