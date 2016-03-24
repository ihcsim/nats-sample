package main

import (
	"fmt"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/nats-io/nats"
)

const msgSubject = "natssample.reply"

func main() {
	natsURL := nats.DefaultURL
	if len(os.Getenv("NATS_HOST")) > 0 {
		natsURL = "nats://" + os.Getenv("NATS_HOST")
	}
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.WithFields(log.Fields{"server-msg": err}).Fatalf("Failed to connect to NATS at %s", natsURL)
	}
	defer nc.Close()

	timeout := 100 * time.Millisecond
	m, err := nc.Request(msgSubject, []byte("Hello World"), timeout)
	if err != nil {
		log.WithFields(log.Fields{"server-msg": err}).Fatal("Request operation failed")
	}

	fmt.Printf("Received reply\nSubject: %s\nData: %s\nReply: %s\n", m.Subject, m.Data, m.Reply)
}
