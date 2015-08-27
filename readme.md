# Siryn
Container Monitoring

Siryn is an opinionated way to add relatively painless monitoring for containers.
It uses [Prometheus](http://prometheus.io) as the collection and alerting
system.  Siryn then pushes stats and configuration for alerting defined by the
container labels.

# Run
To start Siryn, run the following:

```
docker run \
    -p 8080:8080 \
    -p 9090:9090 \
    -p 9091:9091 \
    -d \
    --name siryn \
    -v /var/run/docker.sock:/var/run/docker.sock \
    ehazlett/siryn \
    -D \
    serve
```

# Container Metrics
By default, Siryn will store metrics in Prometheus.  To enable monitoring for
a container, simply add the label `--label siryn=true`.  For example:

```
docker run -d -P --label siryn=true nginx
```

This will start an Nginx container and you should see the stats show up within
ten seconds in the Prometheus UI.

# Prometheus UI
To view the Prometheus admin UI, visit `http://<siryn-host>:9090`.

# NOTE:
This is currently experimental and not yet complete.  Do not use in production.
