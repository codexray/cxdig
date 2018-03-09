package progress

import (
	"codexray/cxdig/core"
	"os"
	"os/signal"

	pb "gopkg.in/cheggaaa/pb.v1"
)

var isMute bool

func SetProgressMuting(val bool) {
	isMute = val
}

// ProgressBar implements core.Progress
type ProgressBar struct {
	impl        *pb.ProgressBar
	isCancelled bool
	cancelFunc  func()
}

func (p *ProgressBar) Init(total int, cancel func()) {
	p.cancelFunc = cancel
	if !isMute {
		p.impl = pb.StartNew(total)
	}
	p.isCancelled = false
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		for sig := range c {
			if sig == os.Interrupt || sig == os.Kill {
				p.isCancelled = true
			}
		}
	}()
}

func (p *ProgressBar) Increment() {
	if p.isCancelled {
		p.Cancel()
	} else {
		if !isMute {
			p.impl.Increment()
		}
	}
}

func (p *ProgressBar) Done() {
	if !isMute {
		p.impl.Finish()
	}
}

func (p *ProgressBar) Cancel() {
	core.Info("Cancelling...")
	p.cancelFunc()
}
