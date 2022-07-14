FROM golang:1.18-alpine AS base
COPY . /src/
WORKDIR /src/example

FROM base AS example-subscriber
RUN go install ./subscriber.go ./common.go ./init.go
CMD ["/go/bin/subscriber"]

FROM base AS example-publisher
RUN go install ./publisher.go ./common.go ./init.go
CMD ["/go/bin/publisher"]
