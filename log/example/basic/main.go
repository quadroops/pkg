package main

import (
	"time"

	"github.com/quadroops/pkg/log"
	"github.com/quadroops/pkg/log/adapter"
)

func main() {

	logger := log.New(adapter.NewZerolog(), "mylogger")
	go func() {
		for i := 0; i < 10; i++ {
			logger.Infof("hello: %d", i)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			logger.Warnf("hello: %d", i)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			logger.Errorf("hello: %d", i)
			time.Sleep(1 * time.Second)
		}
	}()

	logger2 := log.New(adapter.NewZerolog(), "mylogger2",
		log.WithOptionLevel(log.LevelWarn), log.WithAsyncEnabled())

	for i := 0; i < 10; i++ {
		logger2.Debug("This log should not be shown")
		logger2.Info("This log should not be shown")
		logger2.Warn("This log should be shown")
	}

	time.Sleep(10 * time.Second)
}
