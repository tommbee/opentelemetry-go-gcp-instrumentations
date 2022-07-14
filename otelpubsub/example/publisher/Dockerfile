FROM golang:1.18-alpine AS base
COPY . /src/
WORKDIR /src/otelpubsub/example/publisher

FROM base AS example-publisher
RUN go install ./publisher.go
CMD ["/go/bin/publisher"]
