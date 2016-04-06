package main

import (
	"os"
	"os/signal"

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

	msg, quit := make(chan *nats.Msg), make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	nc.Subscribe(msgSubject, func(m *nats.Msg) {
		msg <- m
	})

	for {
		select {
		case m := <-msg:
			log.Infof("Received message\nSubject: %s\nData: %s\nReply: %s\n", m.Subject, m.Data, m.Reply)
			nc.Publish(m.Reply, []byte("Hello there!"))
		case <-quit:
			log.Info("Terminating process...")
			return
		}
	}
}
