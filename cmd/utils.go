package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

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

	path, err := filepath.Abs(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to retrieve absolute path of '%s'", path)
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
