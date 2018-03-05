package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	id := CommitID("testID")
	assert.Equal(t, "testID", id.String())
	id2 := FileType("testID")
	assert.Equal(t, "testID", id2.String())
	id3 := LanguageID("testID")
	assert.Equal(t, "testID", id3.String())
	id4 := ProjectID("testID")
	assert.Equal(t, "testID", id4.String())
	id5 := NormalizeProjectID("testID")
	assert.Equal(t, "testid", id5.String())
}
