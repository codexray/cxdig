package referential

import (
	"sort"

	"github.com/sirupsen/logrus"

	"codexray/cxdig/config"
	"codexray/cxdig/core"
	"codexray/cxdig/types"
)

// LocalDirectoryID identifies a directory belonging to the repository
type LocalDirectoryID int

// Next generates a new file id from the given one
func (id LocalDirectoryID) Next() LocalDirectoryID {
	return id + 1
}

func BuildProjectReferential(commits []types.CommitInfo, registry *config.FileTypeRegistry) types.ProjectReferential {

	core.Info("Building project referential...")

	builder := NewReferentialBuilder(registry)
	for i := len(commits) - 1; i >= 0; i-- {
		commit := commits[i]
		for _, ch := range commit.Changes {
			if ch.FilePath == "" {
				logrus.WithField("commit-id", commit.CommitID).Warn("Ignoring commit change because of empty file path")
				continue
			}
			commitID := commit.CommitID
			authorID := ""
			if commit.Author.Email != "" {
				authorID = commit.Author.Email
			} else {
				authorID = commit.Author.Name
			}
			switch ch.Type {
			case types.FileChangeAdded:
				builder.addFile(ch, commitID, authorID, commit.DateTime)
			case types.FileChangeRenamed:
				builder.renameFile(ch, commitID, authorID, commit.DateTime)
			case types.FileChangeModified:
				builder.modifyFile(ch, commitID, authorID, commit.DateTime)
			case types.FileChangeDeleted:
				builder.deleteFile(ch, authorID, commit.DateTime)
			default:
				logrus.WithFields(logrus.Fields{
					"type":      string(ch.Type),
					"commit-id": commit.CommitID,
					"file":      ch.FilePath,
				}).Panic("Unkown file change type")
			}
		}
	}

	// sort the result to get repeatable output
	result := builder.finalize()
	sort.Sort(result)
	return result
}
