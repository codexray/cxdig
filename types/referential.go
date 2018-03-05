package types

import (
	"time"
)

// LocalFileID identifies a file belonging to the repository
type LocalFileID int

// Next generates a new file id from the given one
func (id LocalFileID) Next() LocalFileID {
	return id + 1
}

// FileNameInfo contains details about the name and path that is or has been used for a file
type FileNameInfo struct {
	FullPath      string    `json:"fullPath,omitempty"`
	EndOfValidity time.Time `json:"endOfValidity,omitempty"`
}

// ActivityInfo is used to know the date and id of a commit that impacted a file
type ActivityInfo struct {
	CommitID CommitID  `json:"commitID,omitempty"`
	Date     time.Time `json:"date,omitempty"`
}

// RenamingInfo contains more details about a commit that renamed a file
type RenamingInfo struct {
	ActivityInfo
	PreviousName string `json:"previousName,omitempty"`
}

// FileActivity contains details about all the events that happened to a file
type FileActivity struct {
	ModificationDates []ActivityInfo `json:"modificationDates,omitempty"`
	RenamingDates     []RenamingInfo `json:"renamingDates,omitempty"`
	RelocationDates   []ActivityInfo `json:"relocationDates,omitempty"`
	UndeletionDates   []ActivityInfo `json:"undeletionDates,omitempty"`
}

// LocalFile contains details about a file that exists or has existed in the repository
type LocalFile struct {
	ID            LocalFileID    `json:"id"`
	LatestPath    string         `json:"latestPath"`
	CreationDate  time.Time      `json:"creationDate"`
	DeletionDate  *time.Time     `json:"deletionDate,omitempty"` // if != nil : file is deleted!
	PreviousNames []FileNameInfo `json:"previousNames,omitempty"`
	FileType      FileType       `json:"fileType,omitempty"`
	Language      LanguageID     `json:"language,omitempty"`
	Activity      FileActivity   `json:"activity,omitempty"`
	AuthorCommits map[string]int `json:"authorCommits,omitempty"`
}

// ProjectReferential details all the files that exist or have existed in a project
type ProjectReferential []LocalFile

// functions used to make it sortable
func (r ProjectReferential) Len() int {
	return len(r)
}

func (r ProjectReferential) Less(i, j int) bool {
	return r[i].ID < r[j].ID
}

func (r ProjectReferential) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
