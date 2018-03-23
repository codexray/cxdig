package cmd

import (
	"codexray/cxdig/core"
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// Values are injected at build time from CI
	softwareVersion string
	buildDate       string
)

func printVersion() {
	if softwareVersion != "" {
		// don't use core.Info() to avoid beinf muted by quiet mode
		fmt.Println(softwareVersion)
	} else {
		core.Warn("version is undefined")
	}
	if buildDate != "" {
		core.Infof("Built on %s with %s\n", buildDate, runtime.Version())
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}
