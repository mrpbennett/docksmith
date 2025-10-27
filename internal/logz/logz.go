// package logger is a helper module for using a custom logger, as i like to have the line number on my logs
// log provided by charmbracelet
package logz

import (
	"os"

	"github.com/charmbracelet/log"
)

func New() *log.Logger {
	return log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
	})
}
