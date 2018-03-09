package progress

import (
	"os"
	"os/signal"
	"syscall"

	pb "github.com/gosuri/uiprogress"
)

var isMute bool

func SetProgressMuting(val bool) {
	isMute = val
}

// ProgressBar Implements core.Progress
type ProgressBar struct {
	Impl        *pb.Bar
	isCancelled bool
}

func (p *ProgressBar) Init(total int) {
	if !isMute {
		pb.Start()
	}
	p.Impl = pb.AddBar(total)
	p.Impl.AppendCompleted()
	p.Impl.PrependElapsed()
	p.isCancelled = false

	// TODO: move that code into a dedicated function
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	go func() {
		<-c
		p.isCancelled = true
	}()
}

func (p *ProgressBar) Increment() {
	p.Impl.Incr()
}

func (p *ProgressBar) Done() {
	//p.Impl.Finish()
}

func (p *ProgressBar) IsCancelled() bool {
	return p.isCancelled
}
