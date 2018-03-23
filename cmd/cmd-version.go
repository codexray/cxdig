package cmd

import (
	"codexray/cxdig/core"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Value is injected at build time
	softwareVersion string
)

func printVersion() {
	if softwareVersion != "" {
		// don't use core.Info() to avoid beinf muted by quiet mode
		fmt.Println(softwareVersion)
	} else {
		core.Warn("Version is undefined!")
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}
