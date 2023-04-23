package cmd

import (
	"strings"

	"github.com/sanderdescamps/mgit/internal/config"
	"github.com/sanderdescamps/mgit/internal/console"
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
			display.Printf("%s", i.GetRepoCloneUrl())
			display.Printf(" path: %s", i.GetFSPath())
			display.Print(" settings:")
			for key, val := range i.GetAllSettings() {
				display.Printf("   - %s: %v", key, val)
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
			display.Printf("%s", repoConf.GetRepoCloneUrl())
			url := repoConf.GetRepoCloneUrl()
			path := repoConf.GetFSPath()
			var subdisplay console.Display
			if Verbose {
				subdisplay = display.GetSubDisplay()
			} else {
				// subdisplay = console.NewNullDisplay()
				subdisplay = display.GetSubDisplay()
			}
			subdisplay.Printf("path: %s", repoConf.GetFSPath())
			if clone, ok := repoConf.GetSettingBool(config.SETTINGS_CLONE_REPO); clone || !ok {
				insecure, ok := repoConf.GetSettingBool("insecure")
				if !ok {
					insecure = false
				}
				repo := repo.NewRepo(url, path, &subdisplay, insecure)
				err := repo.Clone()
				if err != nil {
					display.Final(console.FAILED, err.Error())
				} else {
					display.Final(console.OK, "Repo cloned succesfully")
			}
			} else {
				display.Final(console.SKIPPED, "Repo skipped")
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
			display.Printf("%s", repoConf.GetRepoCloneUrl())
			// display.Printf("path: %s", repoConf.GetFSPath())
			url := repoConf.GetRepoCloneUrl()
			path := repoConf.GetFSPath()
			var subdisplay console.Display
			if Verbose {
				subdisplay = display.GetSubDisplay()
			} else {
				subdisplay = console.NewNullDisplay()
				// subdisplay = display.GetSubDisplay()
			}
			if clone, ok := repoConf.GetSettingBool(config.SETTINGS_CLONE_REPO); clone || !ok {
				insecure, ok := repoConf.GetSettingBool("insecure")
				if !ok {
					insecure = false
				}
				repo := repo.NewRepo(url, path, &subdisplay, insecure)
				if err := repo.CheckTcpConnect(); err != nil {
					display.Errorf("Failed to connect")
				} else if !repo.IsValidRemote() {
					display.Errorf("Invalid remote")
				} else {
					display.PrintColorf("Connection successful", console.GREEN)
				}
			} else {
				display.Final(console.SKIPPED, "Skip repo")
			}
			display.Print(strings.Repeat("-", 80))
		}
	},
}
