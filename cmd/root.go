package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "Upcoming contest CLI app",
		Long:  `We can fetch the details of upcoming codeforces contests`}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
