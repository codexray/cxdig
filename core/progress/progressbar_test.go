package progress

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBar(t *testing.T) {
	p := ProgressBar{}
	p.Init(10)
	p.isCancelled = true
	assert.Equal(t, true, p.IsCancelled())
	p.isCancelled = false
	assert.Equal(t, false, p.IsCancelled())
	for i := 1; i <= 10; i++ {
		p.Increment()
		assert.Equal(t, i, p.Impl.Current())
	}
	p.Done()
}
