package repos

import (
	"codexray/cxdig/core"
	"codexray/cxdig/types"
)

// ProjectName identifies the project being analysed
type ProjectName string

func (p *ProjectName) String() string {
	return string(*p)
}

// Repository interface is used to manipulate a source code repository versioned under a particular CVS
type Repository interface {
	Name() ProjectName
	SampleWithCmd(tool ExternalTool, freq SamplingFreq, limit int, p core.Progress) error
	ExtractCommits() ([]types.CommitInfo, error)
}
