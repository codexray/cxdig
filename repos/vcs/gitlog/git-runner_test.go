package gitlog

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const repoPath = "../../../test_suite/vcs-test/git-repository_test"

func createTestingGitRepo(t *testing.T) {
	cmd := exec.Command("./gitscript-test.sh")
	path, _ := filepath.Abs("../../../test_suite/vcs-test")
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
	path, _ := filepath.Abs("../../../test_suite/vcs-test")
	cmd.Dir = path
	cmd.Run()
}

/*
func TestWalkCommitsWithCommand(t *testing.T) {
	createTestingGitRepo(t)
	repo
	commits, err := ExtractCommitsFromRepository(repoPath)
	assert.NoError(t, err)
	assert.NotPanics(t, func() {
		walkCommitsWithCommand(repoPath, ProjectName("test"), "../../testingApp/testingApp {path}", commits)
	})
	assert.Panics(t, func() {
		WalkCommitsWithCommand(repoPath, ProjectName("test"), "../../testingApp/WrongApp {path}", commits)
	})

	destroyTestingGitRepo(t)
}*/

func TestCheckGitStatus(t *testing.T) {
	createTestingGitRepo(t)

	assert.Equal(t, true, CheckGitStatus(repoPath))

	cmd := exec.Command("touch", "./git-repository_test/testGitStatus.txt")
	path, _ := filepath.Abs("../../../test_suite/vcs-test")
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
