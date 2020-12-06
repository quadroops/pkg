# Log

[![PkgGoDev](https://pkg.go.dev/badge/github.com/quadroops/pkg/log)](https://pkg.go.dev/github.com/quadroops/pkg/log)

It's not a "real" library for logging.  There are so many log library out there, and each of them have their own methods.  
What this library provides are:

- Simple but consistent log interface methods
- Contextual log.  Log based on some context can be an event / a process, or anything, the main point is we can manage our log
more specific and detail.  Contextual log also provide logging metadata and also time measurement.  Metadata will printed out only
if an error happened, using `Warn`.  If you are using contextual logging, all logging's level will be enabled except for `Debug` level.  

We are not try to create a "new log writer", we are using adapter for that purpose.

## Interfaces

This library provide basic methods for logging. 

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
```

## Adapter

Main responsibility of `adapter` is to format the message and the send it to their `writer`.  The format can be anything, json, string, or maybe even a bytes (well i dont know if this idea is exist or not), depends on used adapater.  

Available adapters:

- [zerolog](https://github.com/rs/zerolog)

## Level

You can set minimal log level : 

```go
// adapter using zerolog , this log only enabled if current activity using Info
logger = log.New(adapter.Zerolog(), "mylogger", log.WithOptionLevel(log.LevelInfo))
logger.Info("hello world")
```

Hierarchy :

> Debug -> Info -> Warn -> Error -> Fatal 

Available levels:

```go
log.LevelDebug
log.LevelInfo
log.LevelWarn
log.LevelError
log.LevelFatal
```

## Contextual

The main concept of our "contextual log" is about how we manage / grouping our log based on some specific context.  A "context" can be an event , a process or anything
relate with your application's domain activities.

The differences with "common log" is, contextual used for more detail and specific analytical log.  If you need to analyze your data log, sometimes you need a "context" for that data.

Example:

```go
// process1
logger := log.Contextual(adapter.NewZerolog(), "process1")
logc := logger.Meta(log.KV("requestID", "xid"), log.KV("token", "token"))

// Measure must be called in the end of process, it will calculate time current process
// from the beginning
logc.Info("Incoming request...").Info("Running service logic...").Measure()


// create 'process2' contextual log from previous instance of 'process1' using same adapter
// if you want to use other adapter you'll need to use log.Contextual(...)
logger = logc.NewContextual("process2")
logc = logger.Meta(log.KV("payload", &SomeStruct{}))

// you can get time measurement per log level
logc.Info("Save to database").Measure()
logc.Info("Processing...").Measure()

// do something

// when error happened, all saved metadata will be printed out using Warn
logc.Errorf("Error msg: %s", err.Error()).Measure()
```

An example output:

```
2020-12-05T12:06:01+07:00 INF [process1]: Starting process...
2020-12-05T12:06:01+07:00 INF [process2]: Starting new process...
2020-12-05T12:06:03+07:00 INF [process2]: Process success...
2020-12-05T12:06:03+07:00 INF [process2]: Starting second process...
2020-12-05T12:06:03+07:00 INF [process1]: End porcess...
2020-12-05T12:06:03+07:00 INF [process1]: Measurement: 2.0022325s
2020-12-05T12:06:03+07:00 INF [process1]: Start another process...
2020-12-05T12:06:05+07:00 WRN [process2]: [meta] key1: 1
2020-12-05T12:06:05+07:00 WRN [process2]: [meta] key2: 2
2020-12-05T12:06:05+07:00 WRN [process2]: [meta] key3: 3
2020-12-05T12:06:05+07:00 ERR [process2]: Something went wrong...
2020-12-05T12:06:05+07:00 INF [process2]: Measurement: 4.0050952s
2020-12-05T12:06:08+07:00 WRN [process1]: [meta] requestID: random-id
2020-12-05T12:06:08+07:00 WRN [process1]: [meta] key1: val1
2020-12-05T12:06:08+07:00 WRN [process1]: [meta] key2: 2
2020-12-05T12:06:08+07:00 ERR [process1]: Then error...
2020-12-05T12:06:08+07:00 INF [process1]: Measurement: 7.0059645s
```

## Installation

```
go get -v -u github.com/quadroops/pkg/log
```

## Usages

An example :

```go
// adapter using zerolog 
logger = log.New(adapter.NewZerolog(), "mylogger")
logger.Info("hello world")

// chaining
logger = log.New(adapter.NewZerolog(), "mylogger")
logger.Info("log some variable").Info("another variable").Error("error here")

// adapter using zerolog with additional poster
logger = log.New(adapter.NewZerolog(), "mylogger").WithSender(sender.Console())
logger.Info("hello world")

// without adapter, will use golang log standard library, but, you can't use `Sender`
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
logger = log.New(adapter.Zerolog(), "mylogger", log.WithAsyncEnabled(), log.WithOptionLevel(log.LevelInfo))
logger.Info("hello world") // logging will run in another goroutine 
```
