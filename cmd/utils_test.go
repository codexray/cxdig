package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocateProjectRepositoryPath(t *testing.T) { /*
		id := types.ProjectID("falseProjectID")
		rtn, err := locateProjectRepositoryPath(id)
		assert.Zero(t, rtn)
		assert.Error(t, err)

		id = core.ProjectID("scanner")
		scanProjectID = "../../scanner"
		rtn, err = locateProjectRepositoryPath(id)
		assert.NotZero(t, rtn)
		assert.NoError(t, err)*/
}

func TestDieOnError(t *testing.T) {
	/*err := errors.New("testing error")
	DieOnError(err, "testing msg")
	err = nil
	DieOnError(err, "testing message")*/
}

func TestCheckDirPathExists(t *testing.T) {
	assert.Equal(t, true, checkDirPathExists("./"))
	assert.Equal(t, false, checkDirPathExists("./testingfalsepath"))
}
