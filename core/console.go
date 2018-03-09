package core

import (
	"fmt"
)

var isMute bool

func SetConsoleMuting(val bool) {
	isMute = val
}

func Info(msg string) {
	if !isMute {
		fmt.Println(msg)
	}
}

func Infof(format string, a ...interface{}) {
	if !isMute {
		fmt.Println(fmt.Sprintf(format, a...))
	}
}

func Error(err error) {
	fmt.Println("Error: " + err.Error())
}
