# nats-sample

This project explores the different communication approaches between NATS publishers and subscribers. The code uses the [NATS Go client](http://nats.io/documentation/clients/nats-client-go/) to communicate with a [NATS server Docker container](https://hub.docker.com/_/nats/).

## Getting Started

Start the NATS Docker container:

```sh
$ scripts/start.sh
```

### NATS Request Reply

n a request-response exchange, publish request operation publishes a message with a reply subject expecting a response on that reply subject. The request creates an inbox and performs a request call with the inbox reply and returns the first reply received. Refer hat [NATS documentation](http://nats.io/documentation/concepts/nats-req-rep/) for more information.

To start the request-reply subscriber:

```sh
$ go run cmd/subscriber/reply/main.go
```

When done, use CTL-C to terminate the subscriber.

To start the request-reply publisher:

```sh
$ go run cmd/publisher/reply/main.go
```

By default, the publisher and subscriber looks for the NATS server at `nats://localhost:4222`. Use the `NATS_HOST` environmental variable to override the default server IP address and port.

For example:

```sh
$ NATS_HOST=192.168.99.100:4222 go run publisher/reply/main.go
Received reply
Subject: _INBOX.Y0M7SHOBZ1FHAVJDCE1QH6
Data: Hello there!
Reply:

$ NATS_HOST=192.168.99.100:4222 go run subscriber/reply/main.go
Received message
Subject: natssample.reply
Data: Hello World
Reply: _INBOX.Y0M7SHOBZ1FHAVJDCE1QH6
```

## License

This project is under Apache v2 License. See the [LICENSE file](LICENSE) for the full license text.
