package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleIDString(t *testing.T) {
	id := SampleID(2)
	assert.Equal(t, "2", id.String())
}
