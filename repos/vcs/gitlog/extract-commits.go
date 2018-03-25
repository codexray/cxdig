package gitlog

import (
	"codexray/cxdig/types"
	"strings"

	"github.com/sirupsen/logrus"
)

func ExtractCommitsFromRepository(repoPath string) ([]types.CommitInfo, error) {
	lines, err := RunFullGitLogExtractionWithoutFileDiff(repoPath)
	if err != nil {
		return nil, err
	}
	return parseFullGitLogWithoutDiff(lines), nil
}

func assignCommitNumbers(commits []types.CommitInfo) []types.CommitInfo {
	for i := 0; i < len(commits); i++ {
		reverseNum := len(commits) - i
		commits[i].Number = reverseNum
	}
	return commits
}

func parseFullGitLogWithoutDiff(lines []string) []types.CommitInfo {
	commits := make([]types.CommitInfo, 0, 1000)

	for len(lines) > 0 && lines[0] != "" {
		commit, remaining := extractNextCommitInfo(lines)
		emptyPathFilter := ""
		if commit = filterCommitInfo(commit, emptyPathFilter); commit != nil {
			commits = append(commits, *commit)
		}

		lines = remaining
	}

	return assignCommitNumbers(commits)
}

// filter file changes that do not belong to given project root (if any)
func filterCommitInfo(fullCommit *types.CommitInfo, projectRoot string) *types.CommitInfo {
	if projectRoot == "" {
		return fullCommit
	}

	if !strings.HasSuffix(projectRoot, "/") {
		projectRoot += "/"
	}

	filteredChanges := make([]types.FileChange, 0, len(fullCommit.Changes))
	for _, ch := range fullCommit.Changes {
		keepThisChange := strings.HasPrefix(ch.FilePath, projectRoot)

		// special case: files moved from a filtered location to a non filtered one (or vice versa)
		if ch.Type == types.FileChangeRenamed {
			srcPathIsOk := strings.HasPrefix(ch.FilePath, projectRoot)
			destPathIsOk := strings.HasPrefix(ch.RenamedFile, projectRoot)
			keepThisChange = srcPathIsOk || destPathIsOk

			if keepThisChange && srcPathIsOk != destPathIsOk {
				logrus.WithFields(logrus.Fields{
					"src-path":  ch.FilePath,
					"dest-path": ch.RenamedFile,
				}).Warn("Renamed file partially matches the 'project_root' filter (but file is kept)")
			}
		}

		if keepThisChange {
			// caution: in case of a renamed file, FilePath may not belong to projectRoot
			if strings.HasPrefix(ch.FilePath, projectRoot) {
				ch.FilePath = ch.FilePath[len(projectRoot):]
			}
			if strings.HasPrefix(ch.RenamedFile, projectRoot) {
				ch.RenamedFile = ch.RenamedFile[len(projectRoot):]
			}

			filteredChanges = append(filteredChanges, ch)
		}
	}

	if len(filteredChanges) > 0 {
		result := *fullCommit
		result.Changes = filteredChanges
		return &result
	}

	return nil
}
