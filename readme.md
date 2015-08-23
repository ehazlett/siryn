# Siryn
Container Monitoring

Siryn is an opinionated way to add relatively painless monitoring for containers.
It uses [Prometheus](http://prometheus.io) as the collection and alerting
system.  Siryn then pushes stats and configuration for alerting defined by the
container labels.

# Setup

# Prometheus

## Server

`docker run -d -p 9090:9090 prom/prometheus`

## Push Gateway

`docker run -d -p 9091:9091 prom/pushgateway`
