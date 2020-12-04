package log

import (
	"time"
)

const (
	// LevelDebug .
	LevelDebug = iota

	// LevelInfo .
	LevelInfo

	// LevelWarn .
	LevelWarn

	// LevelError .
	LevelError

	// LevelFatal .
	LevelFatal
)

// Logger is main abstraction
type Logger interface {
	Debug(msg string)
	Debugf(format string, v ...interface{})
	Warn(msg string)
	Warnf(format string, v ...interface{})
	Info(msg string)
	Infof(format string, v ...interface{})
	Error(msg string)
	Errorf(format string, v ...interface{})
	Fatal(msg string)
	Fatalf(format string, v ...interface{})
}

// Sender used as a hook interface methods
type Sender interface {
	Send(msg string)
	Sendf(format string, v ...interface{})
}

// Option used to setup logging level and flag if we need to run it in another goroutine or not
type Option struct {
	Level   int
	IsAsync bool
}

// Log is main object for our logging methods
type Log struct {
	opt     *Option
	adapter Logger
	senders []Sender
	name    string
}

// ContextualLog is main object for contextual logging
type ContextualLog struct {
	adapter   Logger
	startTime time.Time
	meta      []KeyValue
	name      string
}

// Optional used as functional options to our logging
type Optional func(o *Option)

// KeyValue used to save all metadata
type KeyValue struct {
	Key   string
	Value interface{}
}
