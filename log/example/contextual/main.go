package main

import (
	"time"

	"github.com/quadroops/pkg/log"
	"github.com/quadroops/pkg/log/adapter"
)

func processTimes(times int) {
	time.Sleep(time.Duration(times) * time.Second)
}

func main() {
	// process1
	logger := log.Contextual(adapter.NewZerolog(), "process1")
	logc := logger.Meta(log.KV("requestID", "random-id"))

	logc.Info("Starting process...")
	processTimes(2)
	logc.Info("End porcess...")
}
