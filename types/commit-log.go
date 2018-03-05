package types

import (
	"time"
)

// FileChangeType identifies the kind of change done on a file
type FileChangeType string

const (
	// FileChangeAdded identifies a change that added a new file
	FileChangeAdded FileChangeType = "Added"
	// FileChangeDeleted identifies a change that deleted an existing file
	FileChangeDeleted FileChangeType = "Deleted"
	// FileChangeModified identifies a change that modified an existing file
	FileChangeModified FileChangeType = "Modified"
	// FileChangeRenamed identifies a change that renamed an existing file
	FileChangeRenamed FileChangeType = "Renamed"
)

// FileChange contains the details about a change done on a file
type FileChange struct {
	Type        FileChangeType `json:"type"`
	FilePath    string         `json:"file"`
	RenamedFile string         `json:"renamedFile,omitempty"`
}

// AuthorInfo contains details about the author of a commit
type AuthorInfo struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

// CommitID identifies a single commit
type CommitID string

func (id *CommitID) String() string {
	return string(*id)
}

// CommitInfo contains details about a single commit
type CommitInfo struct {
	Number     int          `json:"number"`
	CommitID   CommitID     `json:"commitID"`
	Author     AuthorInfo   `json:"author"`
	DateTime   time.Time    `json:"date"`
	IsMerge    bool         `json:"isMerge,omitempty"`
	Message    string       `json:"message,omitempty"`
	Changes    []FileChange `json:"changes,omitempty"`
	MainParent string       `json:"mainParent,omitempty"`
	Parents    []string     `json:"parents,omitempty"`
}
