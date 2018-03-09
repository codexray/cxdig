package core

type Progress interface {
	Init(total int, cancel func())
	Increment()
	Done()
	Cancel()
}
