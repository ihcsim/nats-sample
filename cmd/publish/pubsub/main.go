package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/nats-io/nats"
)

const msgSubject = "natssample.pubsub"

func main() {
	natsURL := nats.DefaultURL
	if h := os.Getenv("NATS_HOST"); len(h) > 0 {
		natsURL = "nats://" + h
	}

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.WithFields(log.Fields{"server-msg": err}).Fatalf("Failed to connect to NATS at %s", natsURL)
	}
	defer nc.Close()

	msg := "Hello World"
	if err := nc.Publish(msgSubject, []byte(msg)); err != nil {
		log.WithFields(log.Fields{"server-msg": err}).Errorf("Unable to publish message %q\n", msg)
	} else {
		log.Infof("Publishing message %q\n", msg)
	}
}
