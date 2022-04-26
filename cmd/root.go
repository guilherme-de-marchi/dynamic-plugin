package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "bchain-manager",
	Short: "A simple blockchain manager",
}

func Execute() error {
	return rootCmd.Execute()
}

func initConfig() {
	rootCmd.AddCommand()
}

func init() {
	cobra.OnInitialize(initConfig)
}
