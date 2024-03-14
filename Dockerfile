FROM golang:1.21.0-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main ./src/
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080:8080
CMD ["./main"]
