package core

import (
	"fmt"
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

func Warn(msg string) {
	fmt.Println(`/!\ ` + msg)
}
func Warnf(format string, a ...interface{}) {
	fmt.Println(`/!\ ` + fmt.Sprintf(format, a...))
}

// Error reports an error to the user
func Error(err error) {
	fmt.Println("Error: " + err.Error())
}
