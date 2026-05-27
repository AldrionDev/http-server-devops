# -------Build stage-------
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN CGO_ENABLED=0  GOOS=linux go build -o http-server .


# -------Runtime stage-------
FROM alpine:3.23

WORKDIR /app

COPY --from=builder /app/http-server .

EXPOSE 8080

CMD ["./http-server"]