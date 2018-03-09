package gitlog

import (
	"bytes"
	"codexray/cxdig/core"
	"codexray/cxdig/output"
	"codexray/cxdig/repos"
	"codexray/cxdig/types"
	"os"
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

func (r *GitRepository) ConstructSampleList(freq repos.SamplingFreq, commits []types.CommitInfo, limit int, sampleFileName string) error {
	samples := repos.FilterCommitsByStep(commits, freq, limit)
	if len(commits) == 0 {
		logrus.Warn("The filtered list of commits to sample is empty: doing nothing")
		return nil
	}

	core.Info("Sampling repository...")
	if sampleFileName == "" {
		sampleFileName = "samples." + freq.String() + ".json"
	}
	return output.WriteJSONFile(r, sampleFileName, samples)
}

func (r *GitRepository) SampleWithCmd(tool repos.ExternalTool, freq repos.SamplingFreq, commits []types.CommitInfo, sampleFileName string, p core.Progress) error {
	core.Info("Checking repository status...")
	if !CheckGitStatus(r.absPath) {
		return errors.New("the git repository is not clean, commit your changes or track untracked files and retry")
	}
	var samples []types.SampleInfo
	if sampleFileName == "" {
		sampleFileName = "samples." + freq.String() + ".json"
	}
	if err := output.ReadJSONFile(r, sampleFileName, &samples); err != nil {
		return errors.Wrap(err, "failed to load sample file")
	}
	return r.walkCommitsWithCommand(tool, commits, samples, p)
}

func (r *GitRepository) Name() repos.ProjectName {
	name := filepath.Base(r.absPath)
	return repos.ProjectName(name)
}

func (r *GitRepository) walkCommitsWithCommand(tool repos.ExternalTool, commits []types.CommitInfo, samples []types.SampleInfo, p core.Progress) error {
	firstCommitID := commits[0].CommitID.String()

	// TODO: make sure the first commit ID is the current commit ID in the repo
	// restore initial state of the repo
	defer func() {
		core.Info("Restoring original repository state...")
		ResetOnCommit(r.absPath, firstCommitID)
	}()
	core.Info("Executing command on each sample...")
	p.Init(len(samples), func() {
		core.Info("Restoring original repository state...")
		ResetOnCommit(r.absPath, firstCommitID)
		os.Exit(0)
	})
	defer p.Done()

	commitIndex := 0
	treatment := 0
	for _, sample := range samples {
		if p != nil {
			p.Increment()
		}
		for j := commitIndex; j < len(commits); j++ {
			if commits[j].CommitID == sample.CommitID {
				ResetOnCommit(r.absPath, commits[j].CommitID.String())

				cmd := tool.BuildCmd(r.absPath, r.Name(), commits[j])
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
					return errors.Wrap(err, "something wrong happen when running command on commit "+commits[j].CommitID.String())
				}
				logrus.Debug(string(out))
				commitIndex = j
				treatment++
				break
			}
		}
	}
	if treatment != len(samples) {
		logrus.Panic("an error occured while treating samples. All the samples was not treated")
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
