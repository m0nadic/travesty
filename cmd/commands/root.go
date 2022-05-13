package commands

import "github.com/spf13/cobra"

var rootCmd = cobra.Command{
	Use:   "travesty",
	Short: "API mocking tool",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initCmd)
}
