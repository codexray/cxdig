package cmd

import (
	"codexray/cxdig/core"
	"os"

	"github.com/spf13/cobra"
)

var quietMode bool

var rootCmd = &cobra.Command{
	Use:   "scanner",
	Short: "CodeXray tool to scan source code repositories.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		core.SetQuietMode(quietMode)
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
	rootCmd.PersistentFlags().BoolVarP(&quietMode, "quiet", "q", false, "Quiet mode")
}
