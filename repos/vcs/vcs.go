package vcs

import (
	"os"
	"path/filepath"
	"strings"
)

type VcsType string

func (id *VcsType) String() string {
	return string(*id)
}

const (
	GitType     VcsType = "Git"
	GitBareType VcsType = "Git -Bare"
	SvnType     VcsType = "Svn"
	UnknownType VcsType = ""
)

func IdentifyRepositoryType(path string) (VcsType, error) {
	if strings.HasSuffix(path, ".git") {
		return GitBareType, nil
	}

	if dir, err := os.Stat(filepath.Join(path, ".git")); err == nil {
		if dir.IsDir() {
			return GitType, err
		}
	}
	if dir, err := os.Stat(filepath.Join(path, ".svn")); err == nil {
		if dir.IsDir() {
			return SvnType, err
		}
	}

	return UnknownType, nil
}
