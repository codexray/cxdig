package types

import (
	"time"
)

type SampleID int

type SampleInfo struct {
	Number   SampleID  `json:"number"`
	DateTime time.Time `json:"date"`
	CommitID CommitID  `json:"commitID"` // use a pointer to allow null value in JSON output
	AliasOf  SampleID  `json:"aliasOf,omitempty"`
}
