FROM golang:1.25.1 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /blitzdb main.go

FROM alpine:3.18

RUN apk add --no-cache redis

WORKDIR /app

COPY --from=builder /blitzdb /app/blitzdb

RUN ln -s /usr/bin/redis-cli /usr/local/bin/blitzdb-cli

EXPOSE 9999

CMD ["/app/blitzdb"]
