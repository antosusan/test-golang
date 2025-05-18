# Stage 1: Build
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o myapp

# Stage 2: Minimal runtime
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/myapp .

EXPOSE 8080

CMD ["./myapp"]
