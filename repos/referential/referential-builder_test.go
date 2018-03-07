package referential

import (
	"codexray/cxdig/config"
	"codexray/cxdig/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var r = NewReferentialBuilder(config.NewFileTypeRegistry())

func TestAddFile(t *testing.T) {
	ch := types.FileChange{Type: types.FileChangeAdded, FilePath: "test1/path"}
	commitID := types.CommitID("shaCommit1")
	authorID := "author1@mail.com"
	dateTime := time.Now()
	r.addFile(ch, commitID, authorID, dateTime)
	f := r.files[ch.FilePath]
	assert.NotZero(t, f)
	assert.Equal(t, "test1/path", f.LatestPath)
	assert.Equal(t, dateTime, f.CreationDate)
	assert.Equal(t, 1, f.AuthorCommits[authorID])
	assert.Len(t, f.PreviousNames, 0)
	commitID2 := types.CommitID("shaCommit2")
	authorID2 := "author2@mail.com"
	dateTime2 := time.Now().AddDate(0, 0, 1)
	r.addFile(ch, commitID2, authorID2, dateTime2)
	f = r.files[ch.FilePath]
	assert.NotZero(t, f)
	assert.Equal(t, "test1/path", f.LatestPath)
	assert.Equal(t, dateTime, f.CreationDate)
	assert.Equal(t, 1, f.AuthorCommits[authorID])
	assert.Equal(t, 1, f.AuthorCommits[authorID2])
	assert.Len(t, f.PreviousNames, 0)
}

func TestRenameFile(t *testing.T) {
	ch := types.FileChange{Type: types.FileChangeRenamed, FilePath: "test1Modified/path", RenamedFile: "test1Modified/path2"}
	commitID := types.CommitID("shaCommitModified1")
	authorID := "author1@mail.com"
	dateTime := time.Now()
	r.renameFile(ch, commitID, authorID, dateTime)
	f := r.files[ch.FilePath]
	assert.NotZero(t, f)
	assert.Equal(t, "test1Modified/path", f.LatestPath)
	assert.Equal(t, dateTime, f.CreationDate)
	assert.Equal(t, 1, f.AuthorCommits[authorID])
	assert.Len(t, f.PreviousNames, 0)

	ch.RenamedFile = ""
	assert.Panics(t, func() { r.renameFile(ch, commitID, authorID, dateTime) })
	ch.RenamedFile = "test1Modified/path2"

	r = NewReferentialBuilder(config.NewFileTypeRegistry())
	ch.FilePath = ch.RenamedFile
	r.addFile(ch, commitID, authorID, dateTime)
	ch.FilePath = "test1Modified/path"
	r.addFile(ch, commitID, authorID, dateTime)
	r.renameFile(ch, commitID, authorID, dateTime)
	f = r.files[ch.RenamedFile]
	assert.NotZero(t, f)
	assert.Equal(t, "test1Modified/path2", f.LatestPath)
	assert.Equal(t, dateTime, f.CreationDate)
	assert.Equal(t, 2, f.AuthorCommits[authorID])
}

func TestDeleteFile(t *testing.T) {

}
