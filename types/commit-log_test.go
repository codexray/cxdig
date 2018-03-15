package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommitIDString(t *testing.T) {
	id := CommitID("sha1ofarandomcommit")
	assert.Equal(t, "sha1ofarandomcommit", id.String())
}
