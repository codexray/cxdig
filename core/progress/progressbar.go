package progress

import (
	"codexray/cxdig/core"

	pb "gopkg.in/cheggaaa/pb.v1"
)

// ProgressBar implements core.Progress
type ProgressBar struct {
	impl *pb.ProgressBar
}

func (p *ProgressBar) Init(total int) {
	if !core.IsQuietModeEnabled() {
		p.impl = pb.StartNew(total)
	}
}

func (p *ProgressBar) Increment() {
	if !core.IsQuietModeEnabled() {
		p.impl.Increment()
	}
}

func (p *ProgressBar) Done() {
	if !core.IsQuietModeEnabled() {
		p.impl.Finish()
	}
}
