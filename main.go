package main

import (
	"runtime/debug"

	"codexray/cxdig/cmd"

	"github.com/sirupsen/logrus"
)

// HandleCleanExit makes sure to exit properly the application
func HandleCleanExit() {
	if r := recover(); r != nil {
		stackTrace := string(debug.Stack())
		logrus.WithField("stack", stackTrace).Fatalf("PANIC: %v", r)
	}
}

func main() {
	defer HandleCleanExit()

	cmd.Execute()
}
