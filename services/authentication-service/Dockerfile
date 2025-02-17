# Build Stage
FROM golang:latest AS builder

# Pass environment variables as build arguments
ARG PORT

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy application source
COPY . .

# Build the application
RUN go build -o main ./services/authentication-service

# Final Stage
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the built binary from builder
COPY --from=builder /app/main .

# Expose application port (adjust as needed)
EXPOSE $PORT

# Run the executable
ENTRYPOINT ["./main"]
