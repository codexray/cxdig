package cmd

import (
	colorable "github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

func setupLogs(levelStr string) error {
	level, err := logrus.ParseLevel(levelStr)
	if err != nil {
		return err
	}

	logrus.SetLevel(level)

	// make log colors look nice on Windows console
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())
	return nil
}
