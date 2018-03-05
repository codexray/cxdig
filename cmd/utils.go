package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// DieOnError checks for an error and will reports it + die if there is one
func DieOnError(err error, msg string) {
	if err != nil {
		logrus.WithError(err).Error(msg)
		os.Exit(1)
	}
}

func checkDirPathExists(path string) bool {
	if f, err := os.Stat(path); err == nil {
		return f.IsDir()
	}
	return false
}

func getRepositoryPathFromCmdArgs(args []string) (string, error) {
	var path string
	if len(args) == 0 {
		// check if the current dir is a git repository
		var err error
		if path, err = os.Getwd(); err != nil {
			return "", err
		}
	} else if len(args) == 1 {
		path = args[0]
	} else {
		return "", fmt.Errorf("too many arguments")
	}

	// check existence
	if !checkDirPathExists(path) {
		return "", fmt.Errorf("could not find the given path '%s'", path)
	}

	// check it's a git repo
	if !checkDirPathExists(filepath.Join(path, ".git")) {
		return "", fmt.Errorf("'%s' is not a valid git repository", path)
	}

	// handle relative path such as "."
	return filepath.Abs(path)
}
