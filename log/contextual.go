package log

import (
	"fmt"
	"time"
)

// KV used to setup new object KeyValue
func KV(key string, val interface{}) KeyValue {
	return KeyValue{Key: key, Value: val}
}

// Contextual used to create new instance of ContextualLog
func Contextual(adapter Logger, name string) *ContextualLog {
	return &ContextualLog{
		adapter:   adapter,
		name:      name,
		startTime: time.Now().UTC(),
	}
}

// NewContextual used to create new instance from current instance
// but only need to define context name
func (c *ContextualLog) NewContextual(name string) *ContextualLog {
	return Contextual(c.adapter, name)
}

// Meta used to setup metadata contains of key value objects
func (c *ContextualLog) Meta(kv ...KeyValue) *ContextualLog {
	c.meta = append(c.meta, kv...)
	return c
}

// Measure used to setup time measurement.  This method should not return
// current object's instance.
func (c *ContextualLog) Measure() {
	elapsed := time.Since(c.startTime)
	c.adapter.Infof(c.formatMsgWithName("Measurement: %s"), elapsed)
}

// Info used to log something with Info level
func (c *ContextualLog) Info(msg string) *ContextualLog {
	c.adapter.Info(c.formatMsgWithName(msg))
	return c
}

// Infof used to log formatted string with Info level
func (c *ContextualLog) Infof(format string, v ...interface{}) *ContextualLog {
	c.adapter.Infof(c.formatMsgWithName(format), v...)
	return c
}

// Warn used to log something with Warn level
func (c *ContextualLog) Warn(msg string) *ContextualLog {
	c.adapter.Warn(c.formatMsgWithName(msg))
	return c
}

// Warnf used to log formatted string with Warn level
func (c *ContextualLog) Warnf(format string, v ...interface{}) *ContextualLog {
	c.adapter.Warnf(c.formatMsgWithName(format), v...)
	return c
}

// Error used to log something with Error level, and also print out all saved metadata
func (c *ContextualLog) Error(msg string) *ContextualLog {
	c.printOutMeta()
	c.adapter.Error(c.formatMsgWithName(msg))
	return c
}

// Errorf used to log formatted string with Error level, and also print out all saved metadata
func (c *ContextualLog) Errorf(format string, v ...interface{}) *ContextualLog {
	c.printOutMeta()
	c.adapter.Errorf(c.formatMsgWithName(format), v...)
	return c
}

// Fatal used to log something with Fatal level, and also print out all saved metadata
func (c *ContextualLog) Fatal(msg string) *ContextualLog {
	c.printOutMeta()
	c.adapter.Fatal(c.formatMsgWithName(msg))
	return c
}

// Fatalf used to log something with Fatal level, and also print out all saved metadata
func (c *ContextualLog) Fatalf(format string, v ...interface{}) *ContextualLog {
	c.printOutMeta()
	c.adapter.Fatalf(c.formatMsgWithName(format), v...)
	return c
}

func (c *ContextualLog) printOutMeta() {
	for _, kv := range c.meta {
		c.Warnf("[meta] %s: %v", kv.Key, kv.Value)
	}
}

func (c *ContextualLog) formatMsgWithName(msg string) string {
	return fmt.Sprintf("[%s] %s", c.name, msg)
}
