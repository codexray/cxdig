package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var r = ProjectReferential{
	LocalFile{ID: LocalFileID(1)},
	LocalFile{ID: LocalFileID(2)},
	LocalFile{ID: LocalFileID(3)},
}

func TestLen(t *testing.T) {
	assert.Equal(t, 3, r.Len())
}

func TestLess(t *testing.T) {
	assert.True(t, r.Less(0, 1))
	assert.False(t, r.Less(2, 1))
}

func TestSwap(t *testing.T) {
	r.Swap(0, 1)
	assert.Equal(t, LocalFileID(2), r[0].ID)
	assert.Equal(t, LocalFileID(1), r[1].ID)
}

func TestNext(t *testing.T) {
	id := LocalFileID(1)
	id = id.Next()
	assert.Equal(t, LocalFileID(2), id)
}
