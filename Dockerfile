# Start with a lightweight Go base image
FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o myapp .

# Create a minimal final image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myapp .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./myapp"]
