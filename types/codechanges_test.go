package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	diff := LineChanges{
		AddedCount:   11,
		RemovedCount: 3,
		EditedCount:  5,
	}
	diff2 := LineChanges{
		AddedCount:   3,
		RemovedCount: 5,
		EditedCount:  2,
	}
	diff.Merge(diff2)
	assert.Equal(t, 14, diff.AddedCount)
	assert.Equal(t, 8, diff.RemovedCount)
	assert.Equal(t, 7, diff.EditedCount)
}
