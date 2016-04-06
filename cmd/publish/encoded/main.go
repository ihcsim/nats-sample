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

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.WithFields(log.Fields{"server-msg": err}).Fatalf("Failed to establish an encoded connection to NATS at %s", natsURL)
	}
	defer ec.Close()

	send := make(chan *Data)
	defer close(send)
	if err := ec.BindSendChan(msgSubject, send); err != nil {
		log.WithFields(log.Fields{"server-msg": err}).Fatal("Failed to bind send channel")
	}

	d := &Data{Name: "derek", Age: 22, Address: "140 New Montgomery Street"}
	log.Infof("Publishing data %#v\n", d)
	send <- d
}
