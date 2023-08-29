package logger

import "log"

type Logger struct {
	log *log.Logger
}

func New() Logger {
	return Logger{log: log.Default()}
}

func (l Logger) Info(format string, args ...any) {
	l.log.Printf("INFO "+format, args...)
}

func (l Logger) Error(format string, args ...any) {
	l.log.Printf("ERROR "+format, args...)
}
