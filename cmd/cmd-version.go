package cmd

import (
	"codexray/cxdig/core"

	"github.com/spf13/cobra"
)

var (
	// Value is injected at build time
	softwareVersion string
)

func printVersion() {
	if softwareVersion != "" {
		core.Info(softwareVersion)
	} else {
		core.Info("<undefined>")
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}
