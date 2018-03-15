package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguageIDString(t *testing.T) {
	id := LanguageID("golang")
	assert.Equal(t, "golang", id.String())
}

func TestFileTypeString(t *testing.T) {
	ft := FileType("Source")
	assert.Equal(t, "Source", ft.String())
}
