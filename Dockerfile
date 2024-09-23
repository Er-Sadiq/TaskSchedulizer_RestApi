# Use the official Golang image as the builder
FROM golang:1.20 AS builder


WORKDIR /app


COPY go.mod go.sum ./

RUN go mod download


COPY . .

# Build the Go app
RUN go build -o task_schedulizer .


FROM alpine:latest


WORKDIR /app


COPY --from=builder /app/task_schedulizer .


EXPOSE 8081


CMD ["./task_schedulizer"]

