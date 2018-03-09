package progress

import pb "gopkg.in/cheggaaa/pb.v1"

var isMute bool

func SetProgressMuting(val bool) {
	isMute = val
}

// ProgressBar implements core.Progress
type ProgressBar struct {
	impl *pb.ProgressBar
}

func (p *ProgressBar) Init(total int) {
	if !isMute {
		p.impl = pb.StartNew(total)
	}
}

func (p *ProgressBar) Increment() {
	if !isMute {
		p.impl.Increment()
	}
}

func (p *ProgressBar) Done() {
	if !isMute {
		p.impl.Finish()
	}
}
