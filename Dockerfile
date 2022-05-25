FROM golang:1.18-alpine as builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY * ./

RUN go build -o echo .

FROM alpine:latest

WORKDIR /
COPY --from=builder /app/echo .

EXPOSE 8080

USER 1000:1000

ENTRYPOINT ["/echo"]
