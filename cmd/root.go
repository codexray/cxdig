package cmd

import (
	"codexray/cxdig/core"
	"codexray/cxdig/core/progress"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scanner",
	Short: "CodeXray tool to scan source code repositories.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		core.SetConsoleMuting(quiet)
		progress.SetProgressMuting(quiet)
	},
}

var quiet bool

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
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Mute progress bar and information messages, only errors are displayed")
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
