FROM golang:1.6.3-alpine

RUN apk update && apk upgrade && apk add git

RUN mkdir -p /go/src/github.com/byuoitav
ADD . /go/src/github.com/byuoitav/avengineers-slackbot

WORKDIR /go/src/github.com/byuoitav/avengineers-slackbot
RUN go get -d -v
RUN go install -v

CMD ["/go/bin/avengineers-slackbot"]

EXPOSE 9000
