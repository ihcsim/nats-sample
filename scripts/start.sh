#!/bin/bash

NATS_VERSION=${NATS_VERION:-0.7.2}

docker run -d \
  -p 4222:4222 \
  -p 7000:7000 \
  --name nats01 \
  nats:$NATS_VERSION -m 7000
