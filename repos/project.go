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
	// Name returns the (short) name of the repository
	Name() ProjectName
	// GetAbsPath returns the absolute path to the repository root folder
	GetAbsPath() string
	// SampleWithCmd runs a sampling operation with the given tool and sampling list
	SampleWithCmd(tool ExternalTool, rate SamplingRate, commits []types.CommitInfo, samples []types.SampleInfo, p core.Progress) error
	// ExtractCommits extract information about all the commits in the repository
	ExtractCommits() ([]types.CommitInfo, error)
	// HasLocalModifications returns true if the repository has local modification (inclusind non versioned files)
	HasLocalModifications() (bool, error)
}
