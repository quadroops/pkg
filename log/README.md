# Log

It's not a "real" library for logging.  This library is just a simple logging wrapper.  There are so many
log library out there, and each of them has their own methods.  This library is just a simple wrapper but also
provide a simple but consistent interface for logging.

## Interfaces

This library provide basic methods for logging, there are `Logger` and `Sender`. 

```go
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
    AfterLog(msg string)
} 

```

Why just a simple `msg` (string) ? Because this library doesn't want to adding more _complexity_ just to formatting a log message, so it's just a message of string
or a formatted string (actually it's using `fmt.Sprintf`) and that's it.  And actually, you can create your own adapter maybe if you think you need
more custom format like json.

How about `Sender` ? Sometimes, (maybe) we have an activity like : 

- Send a message to apache kafka
- Send a message to elastic search
- Send a message to database
- etc

These activities usually for data/system analytics.  When you are registering a `Sender` object, this method will be executed in a different goroutine.

## Adapter

Main responsibility of `adapter` is to format the message and the send it to their `writer`.  The format can be anything, json, string, or maybe even a bytes (well i dont know if this idea is exist or not).  Available adapters:

- [zerolog](https://github.com/rs/zerolog)

## Sender

For now available sender only a `Console`, this object used just for example.

## Level

You can set minimal log level : 

```go
// adapter using zerolog , this log only enabled if current activity using Info
logger = log.New(adapter.Zerolog(), "mylogger", log.WithOptionLevel(log.LevelInfo))
logger.Info("hello world")
```

Hierarchy :

> Debug -> Warn -> Info -> Error -> Fatal 

Available levels:

```go
log.LevelDebug
log.LevelWarn
log.LevelInfo
log.LevelError
log.LevelFatal
```

## Installation

```
go get -v -u github.com/quadroops/pkg/log
```

## Usages

An example :

```go
// adapter using zerolog 
logger = log.New(adapter.Zerolog(), "mylogger")
logger.Info("hello world")

// adapter using logrus
logger = log.New(adapter.Logrus(), "mylogger")
logger.Info("hello world")

// chaining
logger = log.New(adapter.Logrus(), "mylogger")
logger.Info("log some variable").Info("another variable").Error("error here")

// adapter using zerolog with additional poster
logger = log.New(adapter.Zerolog(), "mylogger").WithSender(sender.Console())
logger.Info("hello world")

// without adapter, will use golang log standard library, if you are prefer this way, you can't use `Sender`
log.Info("hello world")
```

Options available :

```go
type Option struct {
    Level   int     // set minimal log level, by default: LevelDebug
    IsAsync bool    // if this option is enable, all logging method from an adapter will run in another goroutines, by default: false
}
```

Example using options :

```go
// adapter using zerolog , this log only enabled if current activity using Info
logger = log.New(adapter.Zerolog(), "mylogger", log.WithOptionLevel(log.LevelInfo))
logger.Info("hello world")

// set async
logger = log.New(adapter.Zerolog(), "mylogger", log.WithAsyncEnabled())
logger.Info("hello world") // logging will run in another goroutine 

// setup both
logger = log.New(adapter.Zerolog(), "mylogger", log.WithOptionAsyncEnabled(), log.WithOptionLevel(log.LevelInfo))
logger.Info("hello world") // logging will run in another goroutine 
```
