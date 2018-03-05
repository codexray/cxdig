package gitlog

import (
	"codexray/cxdig/types"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type GitLogDecoder struct {
	err error
}

func (d *GitLogDecoder) Err() error {
	return d.err
}

func (d *GitLogDecoder) Clear() {
	d.err = nil
}

func (d *GitLogDecoder) AddError(err string) {
	if d.err == nil {
		d.err = errors.New(err)
	} else {
		d.err = errors.Wrap(d.err, err)
	}
}

func (b *rawCommitBlock) Decode() (*types.CommitInfo, error) {
	decoder := GitLogDecoder{}

	result := &types.CommitInfo{
		Number:   0, // can not be assigned now, to be computed later (in reverse order)
		CommitID: types.CommitID(b.SHA1Text),
		Author:   decoder.DecodeAuthorInfo(b.AuthorText),
		DateTime: decoder.DecodeDateTime(b.DateTimeText),
		IsMerge:  decoder.DecodeMerge(b.MergeText),
		Message:  decoder.DecodeMessage(b.MessageBlock),
		Changes:  decoder.DecodeFileChanges(b.FileChangeBlock),
	}

	if err := decoder.Err(); err != nil {
		return nil, errors.Wrap(err, "failed to decode raw commit log text")
	}
	return result, nil
}

func trimBrackets(text string) string {
	return strings.TrimSpace(strings.TrimPrefix(strings.TrimSuffix(strings.TrimSpace(text), ">"), "<"))
	//return strings.TrimSpace(text[1 : len(text)-1])
}

func (d *GitLogDecoder) DecodeAuthorInfo(text string) types.AuthorInfo {
	if text == "" {
		d.AddError("author information is missing")
		return types.AuthorInfo{}
	}

	result := types.AuthorInfo{}
	if i := strings.Index(text, "<"); i >= 0 {
		result.Name = strings.TrimSpace(text[:i])
		result.Email = trimBrackets(text[i:])
	}

	if result.Email == "" {
		d.AddError("author email is empty")
	}
	if result.Name == "" {
		d.AddError("author name is empty")
	}

	return result
}

func (d *GitLogDecoder) DecodeDateTime(text string) time.Time {
	if text == "" {
		d.AddError("date/time information is missing")
		return time.Time{}
	}

	time, err := time.Parse("Mon, _2 Jan 2006 15:04:05 -0700", text)
	if err != nil {
		d.AddError("failed to parse git log time")
	}

	return time
}

func (d *GitLogDecoder) DecodeMerge(text string) bool {
	return text != "" // for now we ignore the content of the "Merge:" line
}

func (d *GitLogDecoder) DecodeMessage(lines []string) string {
	result := ""
	for _, line := range lines {
		if result != "" {
			result += "\n"
		}
		result += strings.TrimSpace(line)
	}
	return result
}

func decodeSingleFileChange(text string) (*types.FileChange, error) {
	elements := strings.Split(text, "\t")
	if len(elements) < 2 {
		return nil, errors.New("failed to decode file change type")
	}

	tag := elements[0]
	fileName := elements[1]
	renamedName := ""
	var changeType types.FileChangeType

	switch tag[0:1] {
	case "A":
		changeType = types.FileChangeAdded
	case "D":
		changeType = types.FileChangeDeleted
	case "M":
		changeType = types.FileChangeModified
	case "T":
		logrus.WithField("file", fileName).Warn("Treating 'T' file type change as regular file modification (symlink removed?)")
		changeType = types.FileChangeModified
	case "R":
		if len(elements) < 3 {
			return nil, errors.New("failed to decode file change type (renamed name is missing)")
		}
		renamedName = elements[2]
		changeType = types.FileChangeRenamed
	// "C", R", "T", "U":
	default:
		return nil, fmt.Errorf("unkown git change status '%s' in line '%s'", tag, text)
	}

	change := &types.FileChange{
		Type:        changeType,
		FilePath:    fileName,
		RenamedFile: renamedName,
	}
	return change, nil
}

func (d *GitLogDecoder) DecodeFileChanges(lines []string) []types.FileChange {
	changes := make([]types.FileChange, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}

		c, err := decodeSingleFileChange(line)
		if err != nil {
			d.err = errors.Wrap(err, "failed to decode file change")
		} else {
			changes = append(changes, *c)
		}
	}
	return changes
}
