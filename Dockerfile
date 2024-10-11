# Start from the official Golang image
FROM golang:1.16-alpine

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
CMD ["./wait-for-it.sh", "mongo-db:27017", "--", "go", "run", "cmd/api/main.go"]
