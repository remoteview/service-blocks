FROM golang:alpine as builder

ADD . /go/src/github.com/remoteview/service-blocks/

WORKDIR /go/src/github.com/remoteview/service-blocks

RUN set -ex && \
  CGO_ENABLED=0 go build -tags netgo -o service-blocks -v -a -ldflags '-extldflags "-static"' && \
  mv ./service-blocks /usr/bin/service-blocks

FROM busybox
COPY --from=builder /usr/bin/service-blocks /usr/local/bin/service-blocks

EXPOSE 3001

ENTRYPOINT [ "service-blocks" ]
