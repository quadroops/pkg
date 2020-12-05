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
	logc := logger.Meta(log.KV("requestID", "random-id"), log.KV("key1", "val1"), log.KV("key2", 2))

	logc.Info("Starting process...")
	processTimes(2)
	logc.Info("End porcess...").Measure()

	logc.Info("Start another process...")
	processTimes(5)
	logc.Error("Then error...").Measure()
}
