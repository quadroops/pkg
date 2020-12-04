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

	time.Sleep(10 * time.Second)
}
