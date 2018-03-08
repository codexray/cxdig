package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scanner",
	Short: "CodeXray tool to scan source code repositories.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	addCommands()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// addCommands adds child commands to the root command HugoCmd.
func addCommands() {
	// subcommands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(sampleCmd)
}

/*
func init() {
	rootCmd.PersistentFlags().StringP("log-level", "l", "warn", "Level of logs to report")
	cobra.OnInitialize(func() {
		logLevel, _ := rootCmd.PersistentFlags().GetString("log-level")
		if err := setupLogs(logLevel); err != nil {
			core.Error(err)
			os.Exit(1)
		}
	})
}
*/
