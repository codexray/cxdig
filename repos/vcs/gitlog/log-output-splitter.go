package gitlog

import (
	"codexray/cxdig/types"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type rawCommitBlock struct {
	SHA1Text        string
	MergeText       string
	AuthorText      string
	DateTimeText    string
	MessageBlock    []string
	FileChangeBlock []string
}

func extractNextCommitInfo(lines []string) (*types.CommitInfo, []string) {
	if len(lines) == 0 || lines[0] == "" {
		logrus.Warn("Commit log is empty")
		return nil, []string{}
	}
	if !strings.HasPrefix(lines[0], "commit ") {
		logrus.Fatalf("Invalid commit log begining: '%s'", lines[0])
	}

	lineCount := 0
	for _, line := range lines {
		if strings.HasPrefix(line, "commit ") {
			if lineCount > 0 {
				break
			}
		}
		lineCount++
	}

	rawCommitText := lines[:lineCount]
	remaining := lines[lineCount:]

	block, err := newRawCommitBlock(rawCommitText)
	if err != nil {
		panic(err)
	}

	commit, err := block.Decode()
	if err != nil {
		panic(err)
	}

	return commit, remaining
}

func extractValueOfFirstLineWithPrefix(lines []string, prefix string) string {
	for _, line := range lines {
		if strings.HasPrefix(line, prefix) {
			return strings.TrimSpace(line[len(prefix):])
		}
	}
	return ""
}

func newRawCommitBlock(lines []string) (*rawCommitBlock, error) {
	header, remaining := skipUntilFirstEmptyLine(lines)
	commit := &rawCommitBlock{
		SHA1Text:     extractValueOfFirstLineWithPrefix(header, "commit "),
		MergeText:    extractValueOfFirstLineWithPrefix(header, "Merge:"),
		AuthorText:   extractValueOfFirstLineWithPrefix(header, "Author:"),
		DateTimeText: extractValueOfFirstLineWithPrefix(header, "Date:"),
	}

	commit.MessageBlock, remaining = skipUntilFirstEmptyLine(remaining)
	commit.FileChangeBlock, remaining = skipUntilFirstEmptyLine(remaining)
	if len(remaining) > 0 {
		return nil, errors.New("there are some unkown remaining lines")
	}
	return commit, nil
}

func skipUntilFirstEmptyLine(lines []string) ([]string, []string) {
	for i, line := range lines {
		if len(line) == 0 {
			return lines[:i], lines[i+1:]
		}
	}
	return lines, []string{}
}
