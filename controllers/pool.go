package controllers

import (
	"context"
	"golang-event-processor/settings"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func StartPool() {
	workerCount, err := strconv.Atoi(settings.GetSettings().WorkersCount)
	if err != nil {
		log.Printf("WorkerIndex: %d, Type: %T, Error: %v", workerCount, workerCount, err)
	}

	workerConfig := &QueueSettings{
		QueueURL:           settings.GetSettings().QueueUrl,
		MaxNumberOfMessage: 10,
		WaitTimeSecond:     20,
	}

	ctx, cancelFunc := context.WithCancel(context.Background())

	var waitGroup sync.WaitGroup

	for i := 0; i < workerCount; i++ {

		waitGroup.Add(1)
		go worker(ctx, workerConfig, &waitGroup)

	}

	// Handle sigterm and await termChannel signal
	termChannel := make(chan os.Signal)
	signal.Notify(termChannel, syscall.SIGINT, syscall.SIGTERM)

	<-termChannel // Blocks here until interrupted
	cancelFunc()

	waitGroup.Wait()
}
