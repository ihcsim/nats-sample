package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/nats-io/nats"
)

const msgSubject = "natssample.pubsub"
const defaultWorkersCount = 2

func main() {
	natsURL := nats.DefaultURL
	if h := os.Getenv("NATS_HOST"); len(h) > 0 {
		natsURL = "nats://" + h
	}

	workersCount := defaultWorkersCount
	if nw := os.Getenv("WORKERS_COUNT"); len(nw) > 0 {
		var err error
		workersCount, err = strconv.Atoi(nw)
		if err != nil {
			log.Fatalf("Illegal format of variable WORKER_COUNT. Expected a numeric value.")
		}
	}

	// use this slice to keep track of all workers.
	var workers []*Worker
	for i := 0; i < workersCount; i++ {
		w := NewWorker(i)
		workers = append(workers, w)
		go runWorker(w, natsURL)
	}

	// on-sigint, send quit signal to all workers.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	for range quit {
		fmt.Println("Quitting...")
		for _, w := range workers {
			fmt.Println("Sending SIGINT to worker", w.Index)
			w.Close()
			for range w.IsClosed {
				fmt.Println("Closed connection of worker", w.Index)
				break
			}
		}
		return
	}
}

func runWorker(w *Worker, natsURL string) {
	if err := w.Run(msgSubject, natsURL); err != nil {
		log.WithFields(log.Fields{"server-msg": err}).Errorf("Failed to subscribe to NATS server at %s", natsURL)
	}
}
