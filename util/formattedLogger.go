package util

import (
	"fmt"
	"log"
)

type TimeLogger struct {
	*log.Logger
}

func NewTimeLogger(logger *log.Logger) *TimeLogger {
	return &TimeLogger{logger}
}

func (l *TimeLogger) TimePrintf(format string, v ...interface{}) {
	log.Printf("-- %s", fmt.Sprintf(format, v...))
}
