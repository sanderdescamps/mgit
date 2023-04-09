package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/sanderdescamps/mgit/internal/config"
	"github.com/sanderdescamps/mgit/internal/console"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	MGIT_CONFIG_ENV_NAME     = "MGIT_CONFIG_PATH"
	DEFAULT_ROOT_CONFIG_NAME = ".mgit"
)

var Verbose bool
var repoConfigPaths []string
var display console.Display
var repoCfg config.Manager
var rootCfg string

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
	cobra.OnInitialize(initDisplay, initMgitConfig)
	rootCmd.PersistentFlags().StringVarP(&rootCfg, "rconfig", "c", "", fmt.Sprintf("path to root config file (default is $HOME/%s). Not compatible with --repos", DEFAULT_ROOT_CONFIG_NAME))
	rootCmd.PersistentFlags().StringSliceVarP(&repoConfigPaths, "repos", "r", []string{}, "path to repo conf files")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	// rootCmd.AddCommand(cmdClone, cmdPull, cmdConfig)
	rootCmd.AddCommand(cmdTest, cmdConfig, cmdRepo, cmdRepoClone)
	cmdConfig.AddCommand(cmdConfigShow, cmdConfigValidate)
	cmdRepo.AddCommand(cmdRepoClone, cmdRepoCheck)
}

func main() {
	rootCmd.Execute()
}

func initMgitConfig() {
	var mgitConfigPath string
	if rootCfg != "" {
		if fileInfo, err := os.Stat(rootCfg); !errors.Is(err, os.ErrNotExist) && !fileInfo.IsDir() {
			mgitConfigPath = rootCfg
		} else {
			display.Error("root config does not exist")
			os.Exit(1)
		}
	} else if p, ok := os.LookupEnv(MGIT_CONFIG_ENV_NAME); ok {
		if fileInfo, err := os.Stat(p); !errors.Is(err, os.ErrNotExist) && !fileInfo.IsDir() {
			mgitConfigPath = p
		} else {
			display.Errorf("mgit config %s does not exist", p)
			os.Exit(1)
		}

	} else {
		defaultPath := path.Join(myHome(), ".mgit", "config")
		if fileInfo, err := os.Stat(defaultPath); !errors.Is(err, os.ErrNotExist) && !fileInfo.IsDir() {
			mgitConfigPath = defaultPath
		} else {
			display.Warning("No mgit config found, use defaults")
		}
	}

	if mgitConfigPath != "" {
		viper.SetConfigFile(mgitConfigPath)
		viper.SetConfigType("yaml")
		if err := viper.ReadInConfig(); err != nil {
			display.Errorf("Failed to read mgit config %s. (Error:%s)", mgitConfigPath, err.Error())
		}
	}
}

func initDisplay() {
	if Verbose {
		display = *console.NewStdoutDisplay(console.DEBUG)
	} else {
		display = *console.NewStdoutDisplay(console.INFO)
	}
	config.SetDisplay(&display)
}

func loadRepoConfig() {
	repoCfg = config.Manager{}
	viper.AllSettings()

	for _, p := range viper.GetStringSlice("repo_config_paths") {
		repoCfg.LoadRepos(p)
	}
	repoCfg.LoadRepos(repoConfigPaths...)
}
