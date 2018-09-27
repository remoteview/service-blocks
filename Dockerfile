FROM golang:1.11.0-alpine3.8 as builder

ADD . /go/src/github.com/remoteview/service-blocks/

WORKDIR /go/src/github.com/remoteview/service-blocks

RUN set -ex && \
  CGO_ENABLED=0 go build -tags netgo -o service-blocks -v -a -ldflags '-extldflags "-static"' && \
  mv ./service-blocks /usr/bin/service-blocks

# Temporarely using go in order to run migrations on shell.
FROM golang:1.11.0-alpine3.8
COPY --from=builder /usr/bin/service-blocks /usr/local/bin/service-blocks

WORKDIR /service-blocks
COPY --from=builder /go/src/github.com/remoteview/service-blocks/VERSION.txt /service-blocks/VERSION.txt
COPY --from=builder /go/src/github.com/remoteview/service-blocks/database.yml /service-blocks/database.yml
COPY --from=builder /go/src/github.com/remoteview/service-blocks/migrations /service-blocks/migrations

EXPOSE 3001

ENTRYPOINT [ "service-blocks" ]
