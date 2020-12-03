package log

// WithOptionLevel set logging level
func WithOptionLevel(level int) Optional {
	return func(o *Option) {
		o.Level = level
	}
}

// WithAsyncEnabled enable async
func WithAsyncEnabled() Optional {
	return func(o *Option) {
		o.IsAsync = true
	}
}

// New used to create Log instance
func New(adapter Logger, name string, options ...Optional) *Log {
	l := Log{
		adapter: adapter,
		name:    name,
	}

	option := &Option{
		Level:   LevelDebug,
		IsAsync: false,
	}

	for _, of := range options {
		of(option)
	}

	l.opt = option
	return &l
}

// WithSender used to register sender's object to sender registry
func (l *Log) WithSender(sender Sender) *Log {
	l.senders = append(l.senders, sender)
	return l
}

// Debug .
func (l *Log) Debug(msg string) *Log {
	if l.opt.Level > LevelDebug {
		return l
	}

	if l.opt.IsAsync {
		go l.adapter.Debug(msg)
	} else {
		l.adapter.Debug(msg)
	}

	return l
}

// Debugf .
func (l *Log) Debugf(format string, v ...interface{}) *Log {
	if l.opt.Level > LevelDebug {
		return l
	}

	if l.opt.IsAsync {
		go l.adapter.Debugf(format, v...)
	} else {
		l.adapter.Debugf(format, v...)
	}

	return l
}

// Warn .
func (l *Log) Warn(msg string) *Log {
	if l.opt.Level > LevelWarn {
		return l
	}

	if l.opt.IsAsync {
		go l.adapter.Warn(msg)
	} else {
		l.adapter.Warn(msg)
	}

	return l
}

// Warnf .
func (l *Log) Warnf(format string, v ...interface{}) *Log {
	if l.opt.Level > LevelWarn {
		return l
	}

	if l.opt.IsAsync {
		go l.adapter.Warnf(format, v...)
	} else {
		l.adapter.Warnf(format, v...)
	}

	return l
}

// Info .
func (l *Log) Info(msg string) *Log {
	if l.opt.Level > LevelInfo {
		return l
	}

	if l.opt.IsAsync {
		go l.adapter.Info(msg)
	} else {
		l.adapter.Info(msg)
	}

	return l
}

// Infof .
func (l *Log) Infof(format string, v ...interface{}) *Log {
	if l.opt.Level > LevelInfo {
		return l
	}

	if l.opt.IsAsync {
		go l.adapter.Infof(format, v...)
	} else {
		l.adapter.Infof(format, v...)
	}

	return l
}

// Error .
func (l *Log) Error(msg string) *Log {
	if l.opt.Level > LevelError {
		return l
	}

	if l.opt.IsAsync {
		go l.adapter.Error(msg)
	} else {
		l.adapter.Error(msg)
	}

	return l
}

// Errorf .
func (l *Log) Errorf(format string, v ...interface{}) *Log {
	if l.opt.Level > LevelError {
		return l
	}

	if l.opt.IsAsync {
		go l.adapter.Errorf(format, v...)
	} else {
		l.adapter.Errorf(format, v...)
	}

	return l
}

// Fatal .
func (l *Log) Fatal(msg string) *Log {
	if l.opt.IsAsync {
		go l.adapter.Fatal(msg)
	} else {
		l.adapter.Fatal(msg)
	}

	return l
}

// Fatalf .
func (l *Log) Fatalf(format string, v ...interface{}) *Log {
	if l.opt.IsAsync {
		go l.adapter.Fatalf(format, v...)
	} else {
		l.adapter.Fatal(format)
	}

	return l
}
