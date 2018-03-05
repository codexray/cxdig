package vcs

import (
	"codexray/cxdig/repos"
	"codexray/cxdig/repos/vcs/gitlog"
	"errors"
	"fmt"
)

func OpenRepository(path string) (repos.Repository, error) {
	repoType, err := IdentifyRepositoryType(path)
	if repoType == UnknownType || err != nil {
		return nil, fmt.Errorf("the given path is not under a supported version control system")
	}

	switch repoType {
	case GitBareType:
		return nil, errors.New("bare git repositories are not supported")
	case GitType:
		return gitlog.NewGitRepository(path), nil
	case SvnType:
		return nil, errors.New("svn repositories are not supported")
	}

	return nil, nil
}
