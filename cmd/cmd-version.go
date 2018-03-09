package cmd

import (
	"codexray/cxdig/core"
	"codexray/cxdig/core/progress"
	"fmt"
	"time"

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
		pb := &progress.ProgressBar{}
		pb.Init(100) // start rendering

		for i := 0; i < 100; i++ {
			if pb.IsCancelled() {
				break
			}
			pb.Increment()
			time.Sleep(time.Millisecond * 20)
		}
		fmt.Println("hello")
	},
}
