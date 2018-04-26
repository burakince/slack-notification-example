FROM golang:1.10.1-alpine3.7 AS build

RUN apk --no-cache add git
RUN go install github.com/kardianos/govendor && govendor sync
ADD . /go/src/github.com/burakince/slack-notification-example
RUN go install github.com/burakince/slack-notification-example

FROM alpine:3.7

RUN apk --no-cache add --update \
  ca-certificates
RUN mkdir /http-server
WORKDIR /http-server

COPY --from=build /go/bin/slack-notification-example /http-server/slack-notification-example

CMD "/http-server/slack-notification-example"
