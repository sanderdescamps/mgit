package cmd

import (
	"strings"

	"github.com/sanderdescamps/mgit/internal/config"
	"github.com/sanderdescamps/mgit/internal/repo"
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
		loadRepoConfig()
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

var cmdRepoClone = &cobra.Command{
	Use:   "clone",
	Short: "clone repos",
	Long:  `Manage repo configuration (default: show)`,
	PreRun: func(cmd *cobra.Command, args []string) {
		loadRepoConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, repoConf := range repoCfg.GetRepoList() {
			display.Print("%s", repoConf.GetRepoCloneUrl())
			display.Print(" path: %s", repoConf.GetFSPath())
			url := repoConf.GetRepoCloneUrl()
			path := repoConf.GetFSPath()
			if clone, ok := repoConf.GetSettingBool(config.SETTINGS_CLONE_REPO); clone || !ok {
				insecure, ok := repoConf.GetSettingBool("insecure")
				if !ok {
					insecure = false
				}
				repo := repo.NewRepo(url, path, display, insecure)
				repo.Clone()
			}

			display.Print(strings.Repeat("-", 80))
		}
	},
}

var cmdRepoCheck = &cobra.Command{
	Use:   "check",
	Short: "clone repos",
	Long:  `Manage repo configuration (default: show)`,
	PreRun: func(cmd *cobra.Command, args []string) {
		loadRepoConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, repoConf := range repoCfg.GetRepoList() {
			display.Print("%s", repoConf.GetRepoCloneUrl())
			display.Print(" path: %s", repoConf.GetFSPath())
			url := repoConf.GetRepoCloneUrl()
			path := repoConf.GetFSPath()
			if clone, ok := repoConf.GetSettingBool(config.SETTINGS_CLONE_REPO); clone || !ok {
				insecure, ok := repoConf.GetSettingBool("insecure")
				if !ok {
					insecure = false
				}
				repo := repo.NewRepo(url, path, display, insecure)
				if repo.IsValidRemote() {
					display.Okf("repo is reachable")
				} else {
					display.Errorf("unreachable repo")
				}
			}

			display.Print(strings.Repeat("-", 80))
		}
	},
}
