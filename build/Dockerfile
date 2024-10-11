# Go base image to build the app
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o k8s-test .

# Final minimal image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/k8s-test .

# Install necessary certificates and dependencies
RUN apk --no-cache add ca-certificates

# Ensure the binary has execution permissions
RUN chmod +x ./k8s-test

# Set the binary as the entry point
CMD ["./k8s-test"]
