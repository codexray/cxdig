package core

type Progress interface {
	Init(total int)
	Increment()
	Done()
	IsCancelled() bool
}
