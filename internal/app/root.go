package app

import (
	"os"

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

var showVersion bool

func init() {
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Print version")

	cobra.OnInitialize(func() {
		if showVersion {
			PrintVersion()
			os.Exit(0)
		}
		checkForUpdates()
	})
}

func Execute() error {
	return rootCmd.Execute()
}
