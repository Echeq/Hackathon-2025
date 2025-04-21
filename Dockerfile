# Stage 1: Build the application
FROM golang:1.24.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Generate the Kitex code and build the application
RUN kitex -module kitex-multi-protocol idl/user.thrift && \
    go build -o server cmd/server/main.go

# Stage 2: Create a minimal runtime image
FROM alpine:3.18

# Install necessary libraries for running the binary
RUN apk add --no-cache bash

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

# Expose the port the server listens on
EXPOSE 8888

# Command to run the server
CMD ["./server"]