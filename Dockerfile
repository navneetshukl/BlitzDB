
FROM golang:1.25.1 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /blitzdb main.go

FROM alpine:3.18 

WORKDIR /app

COPY --from=builder /blitzdb /app/blitzdb

EXPOSE 8080

CMD ["/app/blitzdb"]