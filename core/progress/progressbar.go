package progress

import pb "gopkg.in/cheggaaa/pb.v1"

// ProgressBar implements core.Progress
type ProgressBar struct {
	impl *pb.ProgressBar
}

func (p *ProgressBar) Init(total int) {
	p.impl = pb.StartNew(total)
}

func (p *ProgressBar) Increment() {
	p.impl.Increment()
}

func (p *ProgressBar) Done() {
	p.impl.Finish()
}
