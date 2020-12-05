package main

import (
	"sync"
	"time"

	"github.com/quadroops/pkg/log"
	"github.com/quadroops/pkg/log/adapter"
)

func processTimes(times int) {
	time.Sleep(time.Duration(times) * time.Second)
}

func process1(logger *log.ContextualLog, wg *sync.WaitGroup) {
	defer wg.Done()

	// process1
	logc := logger.Meta(log.KV("requestID", "random-id"), log.KV("key1", "val1"), log.KV("key2", 2))

	logc.Info("Starting process...")
	processTimes(2)
	logc.Info("End porcess...").Measure()

	logc.Info("Start another process...")
	processTimes(5)
	logc.Error("Then error...").Measure()
}

func process2(logger *log.ContextualLog, wg *sync.WaitGroup) {
	defer wg.Done()

	// process 2
	logger = logger.NewContextual("process2")
	logger.Meta(log.KV("key1", 1), log.KV("key2", 2)).Info("Starting new process...")
	processTimes(2)
	logger.Info("Process success...").Info("Starting second process...")
	processTimes(2)
	logger.Meta(log.KV("key3", 3)).Error("Something went wrong...").Measure()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	logger := log.Contextual(adapter.NewZerolog(), "process1")
	go process1(logger, &wg)
	go process2(logger, &wg)

	wg.Wait()
}
