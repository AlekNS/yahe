FROM golang:1.12 as builder

WORKDIR /home/yaheapi

COPY . .
RUN make build



FROM alpine:3.9

WORKDIR /home/yaheapi
COPY entrypoint.sh /
RUN chmod 755 /entrypoint.sh && adduser -D yaheapi

COPY config.base.yaml /etc/yaheapi/config.yaml

USER yaheapi

COPY --from=builder /home/yaheapi/bin/yaheapi .

EXPOSE 8080
ENTRYPOINT ["/entrypoint.sh", "/home/yaheapi/yaheapi"]
CMD ["serve"]
