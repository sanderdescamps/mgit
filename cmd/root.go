package cmd

import (
	"github.com/sanderdescamps/mgit/internal/config"
	"github.com/sanderdescamps/mgit/internal/console"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Verbose bool
var extraConfigPaths []string
var repoConfigPaths []string
var display console.Display
var repoCfg config.Manager

var rootCmd = &cobra.Command{
	Use: "mgit",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// if Verbose {
		// 	display.SetLogLevel(console.DEBUG)
		// } else {
		// 	display.SetLogLevel(console.INFO)
		// }
	},
}

func init() {
	cobra.OnInitialize(initDisplay, initConfig)
	rootCmd.PersistentFlags().StringSliceVarP(&extraConfigPaths, "config", "c", []string{}, "path to additional config file")
	rootCmd.PersistentFlags().StringSliceVarP(&repoConfigPaths, "repoconf", "r", []string{}, "path to repo conf files")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	// rootCmd.AddCommand(cmdClone, cmdPull, cmdConfig)
	rootCmd.AddCommand(cmdTest, cmdConfig, cmdClone, cmdRepo)
	cmdConfig.AddCommand(cmdConfigShow, cmdConfigValidate)
}

func main() {
	rootCmd.Execute()
}

func initConfig() {
	// cmgr := config.ConfigManager{}
	// cmgr.AddDefaultPaths()

	// //current work dir
	// if wd, err := os.Getwd(); err == nil {
	// 	cmgr.AddConfigPathE(wd, ".mgit", []string{"yml", "yaml", "json"})
	// 	cmgr.AddConfigPathE(wd, "mgit", []string{"yml", "yaml", "json"})
	// } else {
	// 	display.Warningf("failed to get current work directory: %v", err)
	// }

	// if len(extraConfigPaths) > 0 {
	// 	for _, p := range extraConfigPaths {
	// 		if _, err := os.Stat(p); os.IsNotExist(err) {
	// 			display.Warningf("config file (%s) not found", p)
	// 			continue
	// 		}
	// 		cmgr.AddConfigPath(p)
	// 	}
	// }

	// for _, c := range cmgr.GetConfigs() {
	// 	if err := viper.MergeConfigMap(c); err != nil {
	// 		panic(err)
	// 	}
	// }
}

func initDisplay() {
	if Verbose {
		display.SetLogLevel(console.DEBUG)
	} else {
		display.SetLogLevel(console.INFO)
	}
	config.SetDisplay(&display)
}

func initRepo() {
	repoCfg = config.Manager{}

	for _, p := range viper.GetStringSlice("repo_config_paths") {
		repoCfg.LoadRepos(p)
	}
	repoCfg.LoadRepos(repoConfigPaths...)
}
