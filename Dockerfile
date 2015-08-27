FROM alpine:latest
COPY prometheus /bin/prometheus
COPY pushgateway /bin/pushgateway
COPY siryn /bin/siryn
COPY prometheus.yml /opt/siryn/prometheus.yml
WORKDIR /opt/siryn
EXPOSE 8080 9090 9091
ENTRYPOINT ["/bin/siryn"]
CMD ["-h"]
