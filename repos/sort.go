package repos

import "codexray/cxdig/types"

func SortCommitByDateDecr(commits []types.CommitInfo) []types.CommitInfo {
	var tempCommit types.CommitInfo
	for i := 0; i < len(commits); i++ {
		for j := i; j < len(commits); j++ {
			if commits[j].DateTime.After(commits[i].DateTime) {
				tempCommit = commits[i]
				commits[i] = commits[j]
				commits[j] = tempCommit
			}
		}
	}
	return commits
}
