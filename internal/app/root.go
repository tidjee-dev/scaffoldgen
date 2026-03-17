package app

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scaffoldgen",
	Short: "Generate scaffold scripts from markdown structure",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// Check for updates in the background (non-blocking)
		go checkForUpdates()
	},
}

func init() {
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true

	cobra.OnInitialize(func() {
		checkForUpdates()
	})
}

func Execute() error {
	return rootCmd.Execute()
}
