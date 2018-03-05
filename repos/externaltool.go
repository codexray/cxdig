package repos

import (
	"codexray/cxdig/types"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type ExternalTool struct {
	rawCmd string
}

func NewExternalTool(rawCmd string) ExternalTool {
	// caller should validate the given string
	if rawCmd == "" {
		panic("the given command line template is empty")
	}
	return ExternalTool{
		rawCmd: rawCmd,
	}
}

func (tool *ExternalTool) BuildCmd(repoPath string, name ProjectName, commit types.CommitInfo) *exec.Cmd {
	expanded := expandExecRawCmd(tool.rawCmd, repoPath, name, commit)
	toolName, args := splitCommandArgs(expanded)
	return exec.Command(toolName, args...)
}

func expandExecRawCmd(rawcmd string, path string, name ProjectName, commit types.CommitInfo) string {
	rawcmd = strings.Replace(rawcmd, "{path}", path, -1)
	rawcmd = strings.Replace(rawcmd, "{commit.count}", strconv.Itoa(commit.Number), -1)
	rawcmd = strings.Replace(rawcmd, "{commit.id}", commit.CommitID.String(), -1)
	rawcmd = strings.Replace(rawcmd, "{name}", name.String(), -1)

	return rawcmd
}

// TODO: what this function does is not clear at all
func splitCommandArgs(rawcmd string) (string, []string) {
	if strings.Index(rawcmd, "'") >= 0 && strings.Index(rawcmd, `"`) >= 0 {
		logrus.Error(`There is an error with the command syntax (can't have presence of " and ' at the same time) in the --cmd value`)
		os.Exit(1)
	}
	quotePresence := ""
	if strings.Index(rawcmd, "'") >= 0 || strings.Index(rawcmd, `"`) >= 0 {
		if strings.Index(rawcmd, "'") >= 0 {
			quotePresence = "'"
		} else {
			quotePresence = `"`
		}
	}
	TrimCmd := ""
	if quotePresence != "" {
		splittedCmd := strings.Split(rawcmd, quotePresence)
		for i := range splittedCmd {
			if i%2 == 1 {
				splittedCmd[i] = strings.Replace(splittedCmd[i], " ", "/*space*/", -1)
			}
			TrimCmd += splittedCmd[i]
		}
	} else {
		TrimCmd = rawcmd
	}
	rtnCmd := strings.Split(TrimCmd, " ")
	for i := range rtnCmd {
		rtnCmd[i] = strings.Replace(rtnCmd[i], "/*space*/", " ", -1)
	}

	return rtnCmd[0], rtnCmd[1:]
}
