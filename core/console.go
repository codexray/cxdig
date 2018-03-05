package core

import (
	"fmt"
)

func Info(msg string) {
	fmt.Println(msg)
}

func Infof(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf(format, a...))
}

func Error(err error) {
	fmt.Println(err)
}
