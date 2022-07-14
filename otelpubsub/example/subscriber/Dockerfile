FROM golang:1.18-alpine AS base
COPY . /src/
WORKDIR /src/example/subscriber

FROM base AS example-subscriber
RUN go install ./subscriber.go
CMD ["/go/bin/subscriber"]
