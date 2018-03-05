package vcs

/*
func TestIdentifyVCSType(t *testing.T) {
	cmd := exec.Command("./gitscript-test.sh")
	path, _ := filepath.Abs("./vcs-test")
	cmd.Dir = path
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	_, err := cmd.Output()

	errmsg := stderr.String()
	assert.Equal(t, 0, len(errmsg))
	assert.NoError(t, err)

	cmd = exec.Command("./svnscript-test.sh")
	cmd.Dir = path
	cmd.Stderr = &stderr
	_, err = cmd.Output()

	vcsType, err := IdentifyVCSType("./vcs-test/git-repository_test")
	assert.NoError(t, err)
	assert.Equal(t, GitType, vcsType)
	vcsType, err = IdentifyVCSType("./vcs-test/svn-repository_test")
	assert.Equal(t, SvnType, vcsType)

	cmd = exec.Command("rm", "-rf", "./git-repository_test")
	cmd.Dir = path
	cmd.Run()
	cmd = exec.Command("rm", "-rf", "./svn-repository_test")
	cmd.Dir = path
	cmd.Run()
	cmd = exec.Command("rm", "-rf", "./svn-server_test")
	cmd.Dir = path
	cmd.Run()
}
*/
