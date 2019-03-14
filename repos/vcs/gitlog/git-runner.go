package gitlog

import (
	"bytes"
	"codexray/cxdig/types"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// RunFullGitLogExtractionWithoutFileDiff extracts the full git log very quickly but without
// querying the details of the source code changes (which are costly to retrieve)
func RunFullGitLogExtractionWithoutFileDiff(repoPath string) ([]string, error) {
	args := []string{"log", "--name-status", "--date=rfc", "--all"}
	rtn, err := RunGitCommandOnDir(repoPath, args, true)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute git command")
	}
	return rtn, nil
}

// RunSingleCommitDiffExtraction extracts the git log for a single commit with inclusion of the diff
func RunSingleCommitDiffExtraction(repoPath string, commitID types.CommitID) []string {
	args := []string{"log", "-p", "--ignore-all-space", "-n", "1", commitID.String()}
	rtn, err := RunGitCommandOnDir(repoPath, args, true)
	if err != nil {
		logrus.Panic(errors.Wrap(err, "failed to execute git command"))
	}
	return rtn
}

// Set diff.renameLimit to avoid the following error:
// warning: inexact rename detection was skipped due to too many files.
// warning: you may want to set your diff.renameLimit variable to at least 830 and retry the command.
func setGitDiffRenameLimit(repoPath string) {
	cmd := exec.Command("git", "config", "diff.renameLimit", "999999")
	cmd.Dir = repoPath
	if _, err := cmd.Output(); err != nil {
		logrus.WithError(err).Fatal("Failed to configure git repository (diff.renameLimit)")
	}
}

func RunGitCommandOnDir(repoPath string, args []string, setDiff bool) ([]string, error) {
	if setDiff {
		setGitDiffRenameLimit(repoPath)
	}

	cmd := exec.Command("git", args...)
	cmd.Dir = repoPath
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	stdout, err := cmd.Output()
	if errmsg := stderr.String(); len(errmsg) > 0 {
		// TODO: do we need to display those messages?
		// git seems to display messages on stderr for checking operations
		//logrus.WithField("message", errmsg).Warn("git warning/error message")
	}
	if err != nil {
		return nil, err
	}
	return strings.Split(string(stdout), "\n"), nil
}

func CheckOutOnCommit(repoPath string, IDCommit string) ([]string, error) {
	return RunGitCommandOnDir(repoPath, []string{"checkout", "-f", IDCommit}, false)
}

func ClearUntrackedFiles(repoPath string) error {
	_, err := RunGitCommandOnDir(repoPath, []string{"clean", "-fdx"}, false)
	if err != nil {
		return err
	}
	return nil
}

func CheckGitStatus(repoPath string) bool {
	_, err := os.Stat(repoPath)
	if err != nil {
		logrus.Panic("An error occured while trying to check git status of the repo, repository " + repoPath + " not exists")
	}
	out, _ := RunGitCommandOnDir(repoPath, []string{"status", "-s"}, false)
	if len(out) == 1 {
		if out[0] == "" {
			return true
		}
	}
	return false
}

func GetGitCommitsParents(commits []types.CommitInfo, repopath string) []types.CommitInfo {
	args := []string{"rev-list", "--all", "--parents"}
	rtn, _ := RunGitCommandOnDir(repopath, args, true)
	parentMap := make(map[string][]string)
	for _, line := range rtn {
		splittedLine := strings.Split(line, " ")
		if len(splittedLine) > 1 {
			parentMap[splittedLine[0]] = splittedLine[1:]
		}
	}
	for i, commit := range commits {
		if len(parentMap[commit.CommitID.String()]) > 0 {
			commits[i].MainParent = parentMap[commit.CommitID.String()][0]
			commits[i].Parents = parentMap[commit.CommitID.String()]
		}
	}
	return commits
}

func FindAllMergeCommit(repopath string) []string {
	args := []string{"log", "--merges", "--oneline", "--no-abbrev"}
	rtn, _ := RunGitCommandOnDir(repopath, args, true)
	mergeCommits := []string{}
	for _, line := range rtn {
		if line != "" {
			splittedLine := strings.Split(line, " ")
			if len(splittedLine) < 1 {
				logrus.Warning("A commit has no commit message")
			} else {
				mergeCommits = append(mergeCommits, splittedLine[0])
			}
		}
	}
	return mergeCommits
}
