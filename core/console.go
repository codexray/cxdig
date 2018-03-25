package core

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var quietMode bool

// IsQuietModeEnabled returns true if the quiet mode is enabled
func IsQuietModeEnabled() bool {
	return quietMode
}

// SetQuietMode restricts the output messages printed to the user
func SetQuietMode(quiet bool) {
	quietMode = quiet
}

// Info displays a simple message to the user
func Info(msg string) {
	if !quietMode {
		fmt.Println(msg)
	}
}

// Infof displays a formatted message to the user
func Infof(format string, a ...interface{}) {
	if !quietMode {
		fmt.Println(fmt.Sprintf(format, a...))
	}
}

// Warn displays a warning message to the user
func Warn(msg string) {
	color.Yellow("Warning: " + msg)
}

// Warnf displays a fomatted warning message to the user
func Warnf(format string, a ...interface{}) {
	Warn(fmt.Sprintf(format, a...))
}

// Error displays an error message to the user
func Error(err error) {
	color.Red("Error: " + err.Error())
}

// DieOnError check for any error and if so displays it before exiting with an error code
func DieOnError(err error) {
	if err != nil {
		Error(err)
		os.Exit(1)
	}
}
