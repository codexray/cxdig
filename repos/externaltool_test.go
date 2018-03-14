package repos

import (
	"codexray/cxdig/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFindFilterType(t *testing.T) {
	filter, err := DecodeSamplingRate("4c")
	assert.Equal(t, SamplingRate{4, RateCommit}, filter)
	assert.NoError(t, err)
	filter, err = DecodeSamplingRate("1w")
	assert.Equal(t, SamplingRate{1, RateWeek}, filter)
	assert.NoError(t, err)
	filter, err = DecodeSamplingRate("3m")
	assert.Equal(t, SamplingRate{3, RateMonth}, filter)
	assert.NoError(t, err)
	filter, err = DecodeSamplingRate("6q")
	assert.Equal(t, SamplingRate{6, RateQuarter}, filter)
	assert.NoError(t, err)
	filter, err = DecodeSamplingRate("2y")
	assert.Equal(t, SamplingRate{2, RateYear}, filter)
	assert.NoError(t, err)
	filter, err = DecodeSamplingRate("")
	assert.Error(t, err)
	filter, err = DecodeSamplingRate("1")
	assert.Error(t, err)
	filter, err = DecodeSamplingRate("m")
	assert.Error(t, err)
	filter, err = DecodeSamplingRate("cy")
	assert.Error(t, err)
}

/*
func TestFilterCommitInfo(t *testing.T) {
	createTestingGitRepo(t)

	commits, err := ExtractCommitsFromRepository(repoPath)
	assert.NoError(t, err)
	for i, commit := range commits {
		filterCommitInfo(&commit, repoPath)
		assert.Equal(t, commits[i], commit)
	}

	destroyTestingGitRepo(t)
}*/

func TestReplaceRawCmdTemplates(t *testing.T) {
	str := expandExecRawCmd("tool command {path} --name {name}.{sample.number}.{sample.date}.{commit.number}.{commit.id}.{sample.rate}.json --testflag 'with space'",
		"./testPath/testProjet",
		ProjectName("testprojet"),
		types.CommitInfo{Number: 3, CommitID: "testingShaOfCommit"},
		SamplingRate{
			Value: 1,
			Unit:  RateWeek,
		},
		types.SampleInfo{
			Number:   1,
			DateTime: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local),
		})
	assert.Equal(t, "tool command ./testPath/testProjet --name testprojet.1.2000-01-01.3.testingShaOfCommit.1w.json --testflag 'with space'", str)
}

func TestSplitCommandArgs(t *testing.T) {
	toolname, args := splitCommandArgs("tool command ./testPath/testProjet --name testprojet.3.testingShaOfCommit.json --testflag 'with space'")
	assert.Equal(t, "tool", toolname)
	assert.Equal(t, "command", args[0])
	assert.Equal(t, "./testPath/testProjet", args[1])
	assert.Equal(t, "--name", args[2])
	assert.Equal(t, "testprojet.3.testingShaOfCommit.json", args[3])
	assert.Equal(t, "--testflag", args[4])
	assert.Equal(t, "with space", args[5])
	toolname, args = splitCommandArgs(`tool command ./testPath/testProjet --name testprojet.3.testingShaOfCommit.json --testflag "with space"`)
	assert.Equal(t, "tool", toolname)
	assert.Equal(t, "command", args[0])
	assert.Equal(t, "./testPath/testProjet", args[1])
	assert.Equal(t, "--name", args[2])
	assert.Equal(t, "testprojet.3.testingShaOfCommit.json", args[3])
	assert.Equal(t, "--testflag", args[4])
	assert.Equal(t, "with space", args[5])
	toolname, args = splitCommandArgs("tool command ./testPath/testProjet --name testprojet.3.testingShaOfCommit.json --testflag 'withoutspace'")
	assert.Equal(t, "tool", toolname)
	assert.Equal(t, "command", args[0])
	assert.Equal(t, "./testPath/testProjet", args[1])
	assert.Equal(t, "--name", args[2])
	assert.Equal(t, "testprojet.3.testingShaOfCommit.json", args[3])
	assert.Equal(t, "--testflag", args[4])
	assert.Equal(t, "withoutspace", args[5])
	toolname, args = splitCommandArgs("tool command ./testPath/testProjet --name testprojet.3.testingShaOfCommit.json --testflag withoutspace")
	assert.Equal(t, "tool", toolname)
	assert.Equal(t, "command", args[0])
	assert.Equal(t, "./testPath/testProjet", args[1])
	assert.Equal(t, "--name", args[2])
	assert.Equal(t, "testprojet.3.testingShaOfCommit.json", args[3])
	assert.Equal(t, "--testflag", args[4])
	assert.Equal(t, "withoutspace", args[5])
}
