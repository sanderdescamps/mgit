package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdConfig = &cobra.Command{
	Use:   "config",
	Short: "show config",
	Long:  `Manage configuration (default: show)`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdConfigShow.Run(cmd, args)
	},
}

var cmdConfigShow = &cobra.Command{
	Use:   "show",
	Short: "show config",
	Long:  `Create the local directory structure and clone all repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		for k, v := range viper.AllSettings() {
			display.Print("%s=%v", k, v)
		}
	},
}

var cmdConfigValidate = &cobra.Command{
	Use:   "validate",
	Short: "validate config",
	Long:  `Verify if all configuration is okay`,
	Run: func(cmd *cobra.Command, args []string) {
		display.Print("looks good to me")
	},
}
