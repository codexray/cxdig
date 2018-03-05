package repos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFilterType(t *testing.T) {
	filter, err := DecodeSamplingFreq("4c")
	assert.Equal(t, SamplingFreq{4, FreqCommit}, filter)
	assert.NoError(t, err)
	filter, err = DecodeSamplingFreq("1w")
	assert.Equal(t, SamplingFreq{1, FreqWeek}, filter)
	assert.NoError(t, err)
	filter, err = DecodeSamplingFreq("3m")
	assert.Equal(t, SamplingFreq{3, FreqMonth}, filter)
	assert.NoError(t, err)
	filter, err = DecodeSamplingFreq("6q")
	assert.Equal(t, SamplingFreq{6, FreqQuarter}, filter)
	assert.NoError(t, err)
	filter, err = DecodeSamplingFreq("2y")
	assert.Equal(t, SamplingFreq{2, FreqYear}, filter)
	assert.NoError(t, err)
	filter, err = DecodeSamplingFreq("")
	assert.Equal(t, SamplingFreq{1, FreqCommit}, filter)
	assert.NoError(t, err)
	filter, err = DecodeSamplingFreq("1")
	assert.Error(t, err)
	filter, err = DecodeSamplingFreq("m")
	assert.Error(t, err)
	filter, err = DecodeSamplingFreq("cy")
	assert.Error(t, err)
}

func TestFilterCommitInfo(t *testing.T) {
	createTestingGitRepo(t)

	commits, err := ExtractCommitsFromRepository(repoPath)
	assert.NoError(t, err)
	for i, commit := range commits {
		filterCommitInfo(&commit, repoPath)
		assert.Equal(t, commits[i], commit)
	}

	destroyTestingGitRepo(t)
}
