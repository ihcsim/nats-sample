package main

import (
	"os"
	"os/signal"

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

	recv := make(chan *Data)
	if _, err := ec.BindRecvChan(msgSubject, recv); err != nil {
		log.WithFields(log.Fields{"server-msg": err}).Fatal("Failed to bind receive channel")
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
loop:
	for {
		select {
		case d := <-recv:
			log.Infof("Received data %#v\n", d)
		case <-quit:
			log.Infof("Exiting...")
			break loop
		}
	}
}
