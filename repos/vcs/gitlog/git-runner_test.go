package gitlog

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"testing"

	"codexray/cxdig/core"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const repoPath = "../vcs-test/git-repository_test"

func createTestingGitRepo(t *testing.T) {
	cmd := exec.Command("./gitscript-test.sh")
	path, _ := filepath.Abs("../vcs-test")
	cmd.Dir = path
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	_, err := cmd.Output()

	errmsg := stderr.String()
	require.Zero(t, errmsg)
	require.NoError(t, err)
}
func destroyTestingGitRepo(t *testing.T) {
	cmd := exec.Command("rm", "-rf", "./git-repository_test")
	path, _ := filepath.Abs("../vcs-test")
	cmd.Dir = path
	cmd.Run()
}

func TestResetOnCommit(t *testing.T) {
	createTestingGitRepo(t)

	commits, err := ExtractCommitsFromRepository(repoPath)
	assert.NoError(t, err)
	ResetOnCommit(repoPath, commits[1].CommitID.String())
	head := RunGitCommandOnDir(repoPath, []string{"rev-parse", "HEAD"}, false)
	assert.Equal(t, commits[1].CommitID.String(), head[0])

	destroyTestingGitRepo(t)
}

func TestWalkCommitsWithCommand(t *testing.T) {
	createTestingGitRepo(t)

	commits, err := ExtractCommitsFromRepository(repoPath)
	assert.NoError(t, err)
	assert.NotPanics(t, func() {
		WalkCommitsWithCommand(repoPath, ProjectName("test"), "../../testingApp/testingApp {path}", commits)
	})
	assert.Panics(t, func() {
		WalkCommitsWithCommand(repoPath, ProjectName("test"), "../../testingApp/WrongApp {path}", commits)
	})

	destroyTestingGitRepo(t)
}

func TestCheckGitStatus(t *testing.T) {
	createTestingGitRepo(t)

	assert.Equal(t, true, CheckGitStatus(repoPath))

	cmd := exec.Command("touch", "./git-repository_test/testGitStatus.txt")
	path, _ := filepath.Abs("../vcs-test")
	cmd.Dir = path
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	_, err := cmd.Output()

	errmsg := stderr.String()
	assert.Equal(t, 0, len(errmsg))
	assert.NoError(t, err)

	assert.Equal(t, false, CheckGitStatus(repoPath))

	destroyTestingGitRepo(t)
}

func TestGetGitCommitsParents(t *testing.T) {
	createTestingGitRepo(t)

	commits, err := ExtractCommitsFromRepository(repoPath)
	assert.NoError(t, err)
	commits = GetGitCommitsParents(commits, repoPath)
	for i, commit := range commits {
		if i != len(commits)-1 {
			assert.Equal(t, commits[i+1].CommitID.String(), commit.MainParent)
		}
	}

	destroyTestingGitRepo(t)
}

func TestReplaceRawCmdTemplates(t *testing.T) {
	str := expandExecRawCmd("tool command {path} --name {name}.{commit.count}.{commit.id}.json --testflag 'with space'",
		"./testPath/testProjet",
		ProjectName("testprojet"),
		types.CommitInfo{Number: 3, CommitID: "testingShaOfCommit"})
	assert.Equal(t, "tool command ./testPath/testProjet --name testprojet.3.testingShaOfCommit.json --testflag 'with space'", str)
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

func TestFindAllMergeCommit(t *testing.T) {
	createTestingGitRepo(t)
	assert.Len(t, FindAllMergeCommit(repoPath), 0)
	destroyTestingGitRepo(t)
}

func TestFindMainParentOfCommits(t *testing.T) {
	createTestingGitRepo(t)
	commits, _ := ExtractCommitsFromRepository(repoPath)
	commits2 := FindMainParentOfCommits(commits, repoPath)
	for i, _ := range commits {
		assert.Equal(t, commits[i], commits2[i])
	}
	destroyTestingGitRepo(t)
}
