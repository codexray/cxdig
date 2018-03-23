package cmd

import (
	"codexray/cxdig/core"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cxdig",
	Short: "CodeXray tool to scan source code repositories.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// this function is ran in the context of a child command
		// therefore the quiet flag is inherited from its parent and must be
		// checked via Flags() and not PersistentFlags()
		quietMode, err := cmd.Flags().GetBool("quiet")
		if err != nil {
			panic(err)
		}

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
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Quiet mode")
}
