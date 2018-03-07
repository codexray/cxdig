package types

import (
	"time"
)

type SampleInfo struct {
	DateTime time.Time `json:"date"`
	CommitID *CommitID `json:"commitID"` // use a pointer to allow null value in JSON output
}
