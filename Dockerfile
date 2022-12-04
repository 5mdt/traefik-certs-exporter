# syntax=docker/dockerfile:1
## Build
FROM golang:1.19-alpine AS build
WORKDIR /usr/local/go/src/traefik-certs-exporter
COPY *.go ./
RUN go build traefik-certs-exporter

## Deploy
FROM scratch
WORKDIR /
COPY --from=build /usr/local/go/src/traefik-certs-exporter/traefik-certs-exporter /traefik-certs-exporter
ENTRYPOINT ["/traefik-certs-exporter"]
