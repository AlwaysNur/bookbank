# Build stage: use official Go image
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy all source files
COPY . .

# Build the Go application (assumes main.go is your entrypoint)
RUN go build src/main.go

EXPOSE 8080

CMD ["./main"]
