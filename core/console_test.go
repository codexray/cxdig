package core

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuietMode(t *testing.T) {
	SetQuietMode(true)
	assert.Equal(t, true, IsQuietModeEnabled())
	SetQuietMode(false)
	assert.Equal(t, false, IsQuietModeEnabled())
}

func TestInfo(t *testing.T) {
	SetQuietMode(false)
	out := interceptStdOut(func() {
		Info("testing Info")
	})
	assert.Equal(t, "testing Info", strings.TrimRight(out, "\r\n"))

	SetQuietMode(true)
	out = interceptStdOut(func() {
		Info("testing Info")
	})
	assert.Equal(t, "", strings.TrimRight(out, "\r\n"))
}

func TestInfof(t *testing.T) {
	SetQuietMode(false)
	out := interceptStdOut(func() {
		test1 := "Infof"
		test2 := "using testing strings"
		Infof("testing %s while %s", test1, test2)
	})
	assert.Equal(t, "testing Infof while using testing strings", strings.TrimRight(out, "\r\n"))

	SetQuietMode(true)
	out = interceptStdOut(func() {
		test1 := "Infof"
		test2 := "using testing strings"
		Infof("testing %s while %s", test1, test2)
	})
	assert.Equal(t, "", strings.TrimRight(out, "\r\n"))
}

func TestWarn(t *testing.T) {
	SetQuietMode(false)
	out := interceptStdOut(func() {
		Warn("testing Warn")
	})
	assert.Equal(t, `/!\ testing Warn`, strings.TrimRight(out, "\r\n"))

	SetQuietMode(true)
	out = interceptStdOut(func() {
		Warn("testing Warn")
	})
	assert.Equal(t, `/!\ testing Warn`, strings.TrimRight(out, "\r\n"))
}
func TestWarnf(t *testing.T) {
	SetQuietMode(false)
	out := interceptStdOut(func() {
		test1 := "Warnf"
		test2 := "using testing strings"
		Warnf("testing %s while %s", test1, test2)
	})
	assert.Equal(t, `/!\ testing Warnf while using testing strings`, strings.TrimRight(out, "\r\n"))

	SetQuietMode(true)
	out = interceptStdOut(func() {
		test1 := "Warnf"
		test2 := "using testing strings"
		Warnf("testing %s while %s", test1, test2)
	})
	assert.Equal(t, `/!\ testing Warnf while using testing strings`, strings.TrimRight(out, "\r\n"))
}

func TestError(t *testing.T) {
	SetQuietMode(false)
	out := interceptStdOut(func() {
		Error(errors.New("testing Error"))
	})
	assert.Equal(t, "Error: testing Error", strings.TrimRight(out, "\r\n"))

	SetQuietMode(true)
	out = interceptStdOut(func() {
		Error(errors.New("testing Error"))
	})
	assert.Equal(t, "Error: testing Error", strings.TrimRight(out, "\r\n"))
}

func interceptStdOut(printfunc func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	print()
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	printfunc()
	w.Close()
	os.Stdout = old
	out := <-outC
	return out
}
