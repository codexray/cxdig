package gitlog

import (
	"bytes"
	"codexray/cxdig/core"
	"codexray/cxdig/repos"
	"codexray/cxdig/types"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"
)

type GitRepository struct {
	absPath string
}

func NewGitRepository(path string) *GitRepository {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &GitRepository{
		absPath: absPath,
	}
}

func (r *GitRepository) SampleWithCmd(tool repos.ExternalTool, freq repos.SamplingFreq, limit int, p core.Progress) error {
	core.Info("Checking repository status...")
	if !CheckGitStatus(r.absPath) {
		return errors.New("the git repository is not clean, commit your changes and retry")
	}

	core.Info("Scanning repository...")
	commits, err := ExtractCommitsFromRepository(r.absPath)
	if err != nil {
		return err
	}
	commits = repos.FilterCommitsByStep(commits, freq, limit)
	if len(commits) == 0 {
		logrus.Warn("The filtered list of commits to sample is empty: doing nothing")
		return nil
	}

	core.Info("Sampling repository...")
	if tool.IsDefault {
		return r.walkCommitsByDefault(commits, p)
	} else {
		return r.walkCommitsWithCommand(tool, commits, p)
	}
}

func (r *GitRepository) Name() repos.ProjectName {
	name := filepath.Base(r.absPath)
	return repos.ProjectName(name)
}

func (r *GitRepository) walkCommitsByDefault(commits []types.CommitInfo, p core.Progress) error {
	defer func() {
		core.Info("Saving list of commits' sha to treat...")
	}()
	listSha := []types.CommitID{}
	p.Init(len(commits))
	defer p.Done()
	for _, commit := range commits {
		if p != nil {
			p.Increment()
		}
		listSha = append(listSha, commit.CommitID)
	}
	name := r.Name()
	return core.WriteJSONFile(name.String()+".[commitsList].json", listSha)
}

func (r *GitRepository) walkCommitsWithCommand(tool repos.ExternalTool, commits []types.CommitInfo, p core.Progress) error {
	firstCommitID := commits[0].CommitID.String()

	// TODO: make sure the first commit ID is the current commit ID in the repo
	// restore initial state of the repo
	defer func() {
		core.Info("Restoring original repository state...")
		ResetOnCommit(r.absPath, firstCommitID)
	}()

	p.Init(len(commits))
	defer p.Done()

	for _, commit := range commits {
		if p != nil {
			p.Increment()
		}
		ResetOnCommit(r.absPath, commit.CommitID.String())

		cmd := tool.BuildCmd(r.absPath, r.Name(), commit)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		if errmsg := stderr.String(); len(errmsg) > 0 {
			// TODO: better error handling
			//logrus.Warn(errmsg)
		}

		// TODO: evaluate CombinedOutput()
		out, err := cmd.Output()
		if err != nil {
			// TODO: better error message + use defer on ResetOnCommit
			return errors.Wrap(err, "something wrong happen when running command on commit "+commit.CommitID.String())
		}
		logrus.Debug(string(out))
	}
	return nil
}

func (r *GitRepository) ExtractCommits() ([]types.CommitInfo, error) {
	commits, err := ExtractCommitsFromRepository(r.absPath)
	if err != nil {
		return nil, err
	}

	// TODO: check error handling
	commits = GetGitCommitsParents(commits, r.absPath)
	commits = FindMainParentOfCommits(commits, r.absPath)

	return commits, nil
}
