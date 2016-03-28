package main

import (
	"fmt"

	"github.com/nats-io/nats"
)

// Worker encompasses connection, subscription and message handling information of a subscription.
type Worker struct {
	Index    int
	IsClosed chan bool

	conn         *nats.Conn
	subscription *nats.Subscription
	handler      nats.MsgHandler
	quit         chan struct{}
}

// NewWorker initializes a new worker instance with a default message subscription handler.
// The default handler simply prints out the received message.
func NewWorker(index int) *Worker {
	return &Worker{
		Index:    index,
		IsClosed: make(chan bool),

		handler: func(msg *nats.Msg) {
			fmt.Printf("[WORKER %d] Receive a message: %s\n", index, string(msg.Data))
		},
		quit: make(chan struct{}),
	}
}

// Run subscribes the worker to receive subject messages with the NATS server at natsURL.
// It closes the connection when it receives a signal in the quit channel.
func (w *Worker) Run(subject, natsURL string) error {
	if err := w.subscribe(subject, natsURL); err != nil {
		return err
	}

	for range w.quit {
		w.conn.Close()
		w.IsClosed <- w.conn.IsClosed()
		return nil
	}

	return nil
}

// Subscribe sets up the worker to receive subject messages from the NATS server at natsURL.
func (w *Worker) subscribe(subject, natsURL string) error {
	var err error
	if w.conn, err = nats.Connect(natsURL); err != nil {
		return err
	}

	w.subscription, err = w.conn.Subscribe(subject, w.handler)
	if err != nil {
		return err
	}

	return nil
}

// Close tells the worker to close the connection to the NATS server
// by sending a message to the worker's quit channel.
func (w *Worker) Close() {
	w.quit <- struct{}{}
}
