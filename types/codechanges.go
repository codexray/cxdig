package types

// LineChanges contains the details of the changes done in piece of code
type LineChanges struct {
	AddedCount   int `json:"addedCount,omitempty"`
	RemovedCount int `json:"removedCount,omitempty"`
	EditedCount  int `json:"editedCount,omitempty"`
}

// Merge merges another change diff into the current one
func (diff *LineChanges) Merge(other LineChanges) {
	diff.AddedCount += other.AddedCount
	diff.RemovedCount += other.RemovedCount
	diff.EditedCount += other.EditedCount
}

// FileChanges contains details of the changes done in a single file
type FileChanges struct {
	FileName    string      `json:"fileName"`
	LineChanges LineChanges `json:"lineChanges"`
}

// CommitChanges contains details of the changes done by a single commit
type CommitChanges struct {
	CommitID    CommitID      `json:"commitID"`
	FileChanges []FileChanges `json:"fileChanges"`
}
