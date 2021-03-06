# nats-sample

[![Codeship Status for ihcsim/nats-sample](https://app.codeship.com/projects/c64f0310-b667-0134-c27d-5e3aa6825658/status?branch=master)](https://app.codeship.com/projects/194354)

This project explores the different communication approaches between NATS publishers and subscribers. The code uses the [NATS Go client](http://nats.io/documentation/clients/nats-client-go/) to communicate with a [NATS server Docker container](https://hub.docker.com/_/nats/).

## Getting Started

Start the NATS Docker container:

```sh
$ scripts/start.sh
```

### Publish/Subscribe

NATS publish subscribe is a one-to-many communication. A publisher sends a message on a subject. Any active subscriber listening on that subject receives the message. For more information on this toppic, refer the NATS documentation [here](http://nats.io/documentation/concepts/nats-pub-sub/).

To start the pulish-subscribe subscriber:

```sh
$ go run cmd/subscriber/pubsub/main.go cmd/subscriber/pubsub/worker.go
```

When done, use CTL-C to terminate the subscriber.

To start the publish-subscribe publisher:

```sh
$ go run cmd/publisher/pubsub/main.go
```

By default, the publisher and subscriber looks for the NATS server at `nats://localhost:4222`. Use the `NATS_HOST` environmental variable to override the default server IP address and port.

The number of subscriber instances can be modified by setting the `WORKERS_COUNT` environmental variable.

For example:

```sh
$ NATS_HOST=192.168.99.100:4222 go run publisher/pubsub/main.go
Publishing message "Hello World"

$ NATS_HOST=192.168.99.100:4222 WORKERS_COUNT=4 go run publisher/pubsub/main.go publisher/pubsub/worker.go
[WORKER 1] Received a message: Hello World
[WORKER 3] Received a message: Hello World
[WORKER 0] Received a message: Hello World
[WORKER 2] Received a message: Hello World
```

### Request/Reply

In a request-response exchange, publish request operation publishes a message with a reply subject expecting a response on that reply subject. The request creates an inbox and performs a request call with the inbox reply and returns the first reply received. Refer hat [NATS documentation](http://nats.io/documentation/concepts/nats-req-rep/) for more information.

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

### Encoded Channel

An encoded connection wraps a bare connection to a NATS server, providing support for encoding and decoding messages from raw Go types. Refer the NATS [documentation](https://godoc.org/github.com/nats-io/nats#pkg-constants) for supported encoders. This example also demonstrates how to bind channels to encoded communication.

To start the encoded connection subscriber:

```sh
$ go run cmd/subscriber/encoded/main.go
```

When done, use CTL-C to terminate the subscriber.

To start the encoded connection publisher:

```sh
$ go run cmd/publisher/encoded/main.go
```

By default, the publisher and subscriber looks for the NATS server at `nats://localhost:4222`. Use the `NATS_HOST` environmental variable to override the default server IP address and port.

For example:

```sh
$ NATS_HOST=192.168.99.100:4222 go run publisher/encoded/main.go publisher/encoded/data.go
INFO[0000] Publishing data &main.Data{Name:"derek", Age:22, Address:"140 New Montgomery Street"}

$ NATS_HOST=192.168.99.100:4222  go run publisher/encoded/main.go publisher/encoded/data.go
INFO[0004] Received data &main.Data{Name:"derek", Age:22, Address:"140 New Montgomery Street"}
```

## License

This project is under Apache v2 License. See the [LICENSE file](LICENSE) for the full license text.
