# Start from the official Golang image
FROM golang:1.23-alpine AS builder

#Install Bash
RUN apk add --no-cache bash

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Install the wait-for-it script
RUN chmod +x wait-for-it.sh

# Expose the application port
EXPOSE 8080

# Command to run when starting the container
CMD ["go", "run", "cmd/api/main.go"]
